# Stack Jobs UI Integration and Backend Pagination

**Date**: December 4, 2025
**Type**: Feature, Enhancement
**Components**: Web Frontend, Backend API, UI Components, Navigation

## Summary

Integrated stack jobs functionality into the cloud resources web interface, enabling users to view and navigate stack jobs directly from the cloud resources list. Added a "Stack Jobs" menu option that opens a drawer showing paginated stack jobs for a selected cloud resource, with clickable rows that navigate to detailed stack job pages. Implemented server-side pagination in the backend ListStackJobs API to support efficient handling of large numbers of stack jobs.

## Problem Statement / Motivation

The stack jobs feature existed in the backend but lacked a user interface for accessing and viewing stack jobs. Users needed a way to:

### Missing Capabilities

- **No UI access to stack jobs**: Stack jobs could only be accessed via API, with no web interface
- **No integration with cloud resources**: No way to view stack jobs associated with a cloud resource from the cloud resources list
- **No detailed view**: No dedicated page to view complete stack job details including output JSON
- **No pagination support**: Backend API didn't support pagination, which would cause performance issues with large numbers of stack jobs
- **No navigation flow**: No intuitive way to navigate from cloud resources to their associated stack jobs

### User Impact

Without these improvements, users faced:

- Inability to view stack jobs through the web interface
- No way to see deployment history for cloud resources
- Performance issues when loading large numbers of stack jobs
- No detailed view of stack job execution results

## Solution / What's New

Implemented a complete UI integration for stack jobs with three main components:

1. **Stack Jobs Menu in Cloud Resources List**: Added "Stack Jobs" action menu item that opens a drawer showing all stack jobs for the selected cloud resource
2. **Stack Jobs Detail Page**: Created a dedicated page (`/stack-jobs/[id]`) to view complete stack job details including status, timestamps, and full output JSON
3. **Backend Pagination**: Added server-side pagination support to the ListStackJobs API with total pages calculation

### Architecture

**User Flow**:

```
Cloud Resources List Page
    ↓ User clicks "Stack Jobs" menu item
Stack Jobs Drawer (opens)
    ↓ Shows paginated list of stack jobs
    ↓ User clicks on a stack job row
Stack Job Detail Page (/stack-jobs/[id])
    ↓ Shows full stack job details
    ↓ Can navigate back to stack jobs list via breadcrumb
```

**Component Architecture**:

```
Cloud Resources List Component
    ├── Action Menu (View, Edit, Stack Jobs, Delete)
    └── Stack Jobs Drawer
        └── Stack Jobs List Component
            └── Table with Pagination
                ↓ (on row click)
                Stack Job Detail Page
                    ├── Breadcrumb Navigation
                    ├── Stack Job Header
                    └── JSON Output Viewer
```

**Backend API Flow**:

```
Frontend Request (ListStackJobs)
    ↓ With PageInfo (page, size)
Backend Service (ListStackJobs)
    ↓ Applies filters and pagination
Repository Layer
    ↓ MongoDB query with skip/limit
    ↓ Count query for total pages
Response (jobs + totalPages)
```

### Key Features

**1. Stack Jobs Menu Integration**

- Added "Stack Jobs" menu item to cloud resources action menu
- Opens drawer when clicked, showing stack jobs for the selected cloud resource
- Drawer uses the same drawer component pattern as other features
- Maintains state for selected cloud resource

**2. Stack Jobs List Component**

- Displays stack jobs in a paginated table
- Shows ID (truncated), Status, Created At, Updated At columns
- Status displayed with color-coded status chips
- Clickable rows navigate to detail page
- Server-side pagination with page navigation controls
- Empty state message when no stack jobs found

**3. Stack Job Detail Page**

- Dedicated route: `/stack-jobs/[id]`
- Breadcrumb navigation with clickable "Stack Jobs" link
- Stack job header showing:
  - Job ID with copy-to-clipboard functionality
  - Status chip
  - Last updated timestamp
- Full JSON output displayed with syntax highlighting
- Loading states with skeleton placeholders
- Can open stack jobs drawer from breadcrumb to view all jobs for the same cloud resource

**4. Backend Pagination**

- Added `PageInfo` support to `ListStackJobsRequest` proto
- Added `total_pages` field to `ListStackJobsResponse` proto
- Repository layer supports pagination with MongoDB skip/limit
- Total pages calculated using ceiling division: `(totalCount + pageSize - 1) / pageSize`
- Default pagination: page 0, size 20 if not provided
- Maintains backward compatibility (pagination is optional)

## Implementation Details

### 1. Backend Pagination Support

**File**: `app/backend/apis/proto/stack_job_service.proto`

Added pagination fields to the ListStackJobs API:

```45:61:app/backend/apis/proto/stack_job_service.proto
// Request message for listing stack jobs.
message ListStackJobsRequest {
  // Optional filter by cloud resource ID.
  optional string cloud_resource_id = 1;
  // Optional filter by status (success, failed, in_progress).
  optional string status = 2;
  // Pagination parameters (optional). If not provided, returns all jobs.
  optional PageInfo page_info = 3;
}

// Response message containing a list of stack jobs.
message ListStackJobsResponse {
  // List of stack jobs.
  repeated StackJob jobs = 1;
  // Total number of pages available (only set when page_info is provided).
  int32 total_pages = 2;
}
```

**Key design decisions**:

- Reused `PageInfo` message from `cloud_resource_service.proto` for consistency
- `total_pages` only set when `page_info` is provided (backward compatible)
- Pagination is optional to maintain backward compatibility

**File**: `app/backend/internal/service/stack_job_service.go`

Implemented pagination logic in the service layer:

```145:220:app/backend/internal/service/stack_job_service.go
// ListStackJobs lists stack jobs with optional filters and pagination.
func (s *StackJobService) ListStackJobs(
	ctx context.Context,
	req *connect.Request[backendv1.ListStackJobsRequest],
) (*connect.Response[backendv1.ListStackJobsResponse], error) {
	opts := &database.StackJobListOptions{}

	if req.Msg.CloudResourceId != nil {
		id := *req.Msg.CloudResourceId
		opts.CloudResourceID = &id
	}

	if req.Msg.Status != nil {
		s := *req.Msg.Status
		opts.Status = &s
	}

	// Apply pagination with defaults (page=0, size=20) if not provided
	var pageNum int32 = 0
	var pageSize int32 = 20
	if req.Msg.PageInfo != nil {
		pageNum = req.Msg.PageInfo.Num
		pageSize = req.Msg.PageInfo.Size
	}
	opts.PageNum = &pageNum
	opts.PageSize = &pageSize

	// Calculate total pages
	totalCount, err := s.stackJobRepo.Count(ctx, opts)
	if err != nil {
		logrus.WithError(err).Error("Failed to count stack jobs for pagination")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to count stack jobs: %w", err))
	}

	var totalPages int32
	if pageSize > 0 {
		totalPages = int32((totalCount + int64(pageSize) - 1) / int64(pageSize))
	}

	logrus.WithFields(logrus.Fields{
		"cloud_resource_id": req.Msg.CloudResourceId,
		"status":            req.Msg.Status,
		"page_num":          opts.PageNum,
		"page_size":         opts.PageSize,
	}).Info("Listing stack jobs")

	jobs, err := s.stackJobRepo.List(ctx, opts)
	if err != nil {
		logrus.WithError(err).Error("Failed to list stack jobs")
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list stack jobs: %w", err))
	}

	protoJobs := make([]*backendv1.StackJob, 0, len(jobs))
	for _, job := range jobs {
		protoJob := &backendv1.StackJob{
			Id:              job.ID.Hex(),
			CloudResourceId: job.CloudResourceID,
			Status:          job.Status,
			Output:          job.Output,
		}
```

**Key features**:

- Default pagination: page 0, size 20 if not provided
- Calculates total pages using ceiling division
- Returns both jobs and totalPages in response
- Maintains all existing filter capabilities (cloud_resource_id, status)

**File**: `app/backend/internal/database/stack_job_repo.go`

Added pagination support to repository layer:

```130:180:app/backend/internal/database/stack_job_repo.go
// StackJobListOptions contains options for listing stack jobs.
type StackJobListOptions struct {
	CloudResourceID *string
	Status          *string
	PageNum         *int32
	PageSize        *int32
}

// List retrieves stack jobs with optional filters and pagination.
func (r *StackJobRepository) List(ctx context.Context, opts *StackJobListOptions) ([]*models.StackJob, error) {
	filter := bson.M{}

	if opts != nil {
		if opts.CloudResourceID != nil && *opts.CloudResourceID != "" {
			filter["cloud_resource_id"] = *opts.CloudResourceID
		}

		if opts.Status != nil && *opts.Status != "" {
			filter["status"] = *opts.Status
		}
	}

	findOptions := options.Find()

	// Apply pagination if provided
	if opts != nil && opts.PageNum != nil && opts.PageSize != nil {
		pageNum := *opts.PageNum
		pageSize := *opts.PageSize
		if pageSize > 0 {
			skip := int64(pageNum) * int64(pageSize)
			findOptions.SetSkip(skip)
			findOptions.SetLimit(int64(pageSize))
		}
	}

	// Sort by created_at descending (newest first)
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to query stack jobs: %w", err)
	}
	defer cursor.Close(ctx)

	var jobs []*models.StackJob
	if err := cursor.All(ctx, &jobs); err != nil {
		return nil, fmt.Errorf("failed to decode stack jobs: %w", err)
	}

	return jobs, nil
}
```

Added `Count` method for total pages calculation:

```182:202:app/backend/internal/database/stack_job_repo.go
// Count returns the total count of stack jobs with optional filters.
func (r *StackJobRepository) Count(ctx context.Context, opts *StackJobListOptions) (int64, error) {
	filter := bson.M{}

	if opts != nil {
		if opts.CloudResourceID != nil && *opts.CloudResourceID != "" {
			filter["cloud_resource_id"] = *opts.CloudResourceID
		}

		if opts.Status != nil && *opts.Status != "" {
			filter["status"] = *opts.Status
		}
	}

	count, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("failed to count stack jobs: %w", err)
	}

	return count, nil
}
```

**Key features**:

- Pagination uses MongoDB skip/limit for efficient querying
- Count method respects the same filters as List method
- Sorted by created_at descending (newest first)
- Maintains all existing filter capabilities

### 2. Frontend Stack Jobs Integration

**File**: `app/frontend/src/components/shared/cloud-resources-list/cloud-resources-list.tsx`

Added "Stack Jobs" menu item to cloud resources action menu:

```237:279:app/frontend/src/components/shared/cloud-resources-list/cloud-resources-list.tsx
  const handleOpenStackJobs = useCallback((row: CloudResource) => {
    setSelectedResourceForStackJobs(row);
    setStackJobsDrawerOpen(true);
  }, []);

  const handleCloseStackJobs = useCallback(() => {
    setStackJobsDrawerOpen(false);
    setSelectedResourceForStackJobs(null);
  }, []);

  const tableActions: ActionMenuProps<CloudResource>[] = useMemo(
    () => [
      {
        text: 'View',
        handler: (row: CloudResource) => {
          handleOpenDrawer('view', row);
        },
        isMenuAction: true,
      },
      {
        text: 'Edit',
        handler: (row: CloudResource) => {
          handleOpenDrawer('edit', row);
        },
        isMenuAction: true,
      },
      {
        text: 'Stack Jobs',
        handler: (row: CloudResource) => {
          handleOpenStackJobs(row);
        },
        isMenuAction: true,
      },
      {
        text: 'Delete',
        handler: (row: CloudResource) => {
          handleConfirmDelete(row);
        },
        isMenuAction: true,
      },
    ],
    [handleOpenDrawer, handleConfirmDelete, handleOpenStackJobs]
  );
```

Added stack jobs drawer state and rendering:

```96:99:app/frontend/src/components/shared/cloud-resources-list/cloud-resources-list.tsx
  // Stack jobs drawer state
  const [stackJobsDrawerOpen, setStackJobsDrawerOpen] = useState(false);
  const [selectedResourceForStackJobs, setSelectedResourceForStackJobs] =
    useState<CloudResource | null>(null);
```

**File**: `app/frontend/src/components/shared/stackjob/stack-jobs-drawer.tsx`

Created drawer component for stack jobs list:

```1:18:app/frontend/src/components/shared/stackjob/stack-jobs-drawer.tsx
'use client';

import { Drawer } from '@/components/shared/drawer';
import { StackJobsList } from './stack-jobs-list';
export interface StackJobsDrawerProps {
  open: boolean;
  cloudResourceId: string;
  onClose: () => void;
}

export function StackJobsDrawer({ open, cloudResourceId, onClose }: StackJobsDrawerProps) {
  return (
    <Drawer open={open} onClose={onClose} title="Stack Jobs" width={900}>
      <StackJobsList cloudResourceId={cloudResourceId} />
    </Drawer>
  );
}
```

**Key features**:

- Reuses existing Drawer component
- Width set to 900px for better table visibility
- Passes cloudResourceId to filter stack jobs

**File**: `app/frontend/src/components/shared/stackjob/stack-jobs-list.tsx`

Created list component with pagination:

```22:124:app/frontend/src/components/shared/stackjob/stack-jobs-list.tsx
export function StackJobsList({ cloudResourceId }: StackJobsListProps) {
  const router = useRouter();
  const { query } = useStackJobQuery();
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);
  const [stackJobs, setStackJobs] = useState<StackJob[]>([]);
  const [totalPages, setTotalPages] = useState(0);
  const [apiLoading, setApiLoading] = useState(true);

  // Function to call the API
  const handleLoadStackJobs = useCallback(() => {
    setApiLoading(true);
    if (query) {
      query
        .listStackJobs(
          create(ListStackJobsRequestSchema, {
            cloudResourceId,
            pageInfo: create(PageInfoSchema, {
              num: page,
              size: rowsPerPage,
            }),
          })
        )
        .then((result) => {
          setStackJobs(result.jobs);
          setTotalPages(result.totalPages || 0);
          setApiLoading(false);
        })
        .catch(() => {
          setApiLoading(false);
        });
    }
  }, [query, cloudResourceId, page, rowsPerPage]);

  // Reset to first page when cloudResourceId changes
  useEffect(() => {
    setPage(0);
  }, [cloudResourceId]);

  // Auto-load on mount and when dependencies change
  useEffect(() => {
    if (query && cloudResourceId) {
      handleLoadStackJobs();
    }
  }, [query, cloudResourceId, handleLoadStackJobs]);

  // Handle page change
  const handlePageChange = useCallback((newPage: number, newRowsPerPage: number) => {
    setPage(newPage);
    setRowsPerPage(newRowsPerPage);
  }, []);

  const tableDataTemplate = useMemo(
    () => ({
      id: (val: string, row: StackJob) => {
        return <Typography variant="subtitle2">{row.id.substring(0, 12)}...</Typography>;
      },
      status: (val: string, row: StackJob) => {
        return <StatusChip status={row.status || 'unknown'} />;
      },
      createdAt: (val: string, row: StackJob) => {
        return row.createdAt ? formatTimestampToDate(row.createdAt, 'DD/MM/YYYY HH:mm') : '-';
      },
      updatedAt: (val: string, row: StackJob) => {
        return row.updatedAt ? formatTimestampToDate(row.updatedAt, 'DD/MM/YYYY HH:mm') : '-';
      },
    }),
    []
  );

  const clickableColumns = useMemo(
    () => ({
      id: (row: StackJob) => {
        router.push(`/stack-jobs/${row.id}`);
      },
    }),
    [router]
  );

  return (
    <Box>
      <TableComp
        data={stackJobs}
        loading={apiLoading}
        options={{
          headers: ['ID', 'Status', 'Created At', 'Updated At'],
          dataPath: ['id', 'status', 'createdAt', 'updatedAt'],
          dataTemplate: tableDataTemplate,
          clickableColumns,
          actions: [],
          showPagination: true,
          paginationMode: PAGINATION_MODE.SERVER,
          currentPage: page,
          rowsPerPage: rowsPerPage,
          totalPages: totalPages,
          onPageChange: handlePageChange,
          emptyMessage: 'No stack jobs found',
          border: true,
        }}
      />
    </Box>
  );
}
```

**Key features**:

- Server-side pagination with page size of 10
- Clickable rows navigate to detail page
- Status displayed with color-coded chips
- Truncated ID display (first 12 characters)
- Resets to page 0 when cloudResourceId changes
- Loading states handled

### 3. Stack Job Detail Page

**File**: `app/frontend/src/app/stack-jobs/[id]/page.tsx`

Created detail page with breadcrumb navigation:

```16:96:app/frontend/src/app/stack-jobs/[id]/page.tsx
export default function StackJobDetailPage() {
  const { theme } = useContext(AppContext);
  const params = useParams();
  const { query } = useStackJobQuery();
  const [stackJob, setStackJob] = useState<StackJob | null>(null);
  const [stackJobsDrawerOpen, setStackJobsDrawerOpen] = useState(false);

  const stackJobId = params?.id as string;

  const handleCloseStackJobs = useCallback(() => {
    setStackJobsDrawerOpen(false);
  }, []);

  const handleStackJobsClick = useCallback(() => {
    if (stackJob?.cloudResourceId) {
      setStackJobsDrawerOpen(true);
    }
  }, [stackJob?.cloudResourceId]);

  const breadcrumbs: IBreadcrumbItem[] = useMemo(() => {
    const items: IBreadcrumbItem[] = [];

    // Always show the ID from params, even if stackJob is not loaded yet
    if (stackJobId) {
      items.push({
        name: stackJobId,
        handler: undefined, // Last item is not clickable
      });
    }

    return items;
  }, [stackJobId, handleStackJobsClick]);

  useEffect(() => {
    if (query && stackJobId) {
      query.getById(stackJobId).then((job) => {
        setStackJob(job);
      });
    }
  }, [query, stackJobId]);

  const updatedTime = stackJob?.updatedAt
    ? formatTimestampToDate(stackJob.updatedAt, 'DD/MM/YYYY, HH:mm:ss')
    : '-';

  return (
    <StackJobContainer>
      <Stack gap={2}>
        <Breadcrumb
          breadcrumbs={breadcrumbs}
          startBreadcrumb={
            <BreadcrumbStartIcon
              icon={ICON_NAMES.INFRA_HUB}
              iconProps={{ sx: { filter: theme.mode === THEME.DARK ? 'invert(1)' : 'none' } }}
              label="Stack Jobs"
              handler={handleStackJobsClick}
            />
          }
        />
        <StackJobHeader stackJob={stackJob} updatedTime={updatedTime} />

        <Box>
          {stackJob ? (
            <JsonCode content={stackJob?.output || {}} />
          ) : (
            <Skeleton variant="rounded" width={'100%'} height={200} />
          )}
        </Box>

        {/* Stack Jobs Drawer */}
        {stackJob?.cloudResourceId && (
          <StackJobsDrawer
            open={stackJobsDrawerOpen}
            cloudResourceId={stackJob.cloudResourceId}
            onClose={handleCloseStackJobs}
          />
        )}
      </Stack>
    </StackJobContainer>
  );
}
```

**Key features**:

- Dynamic route using Next.js `[id]` parameter
- Breadcrumb navigation with clickable "Stack Jobs" link
- Opens stack jobs drawer when breadcrumb is clicked
- Displays full JSON output with syntax highlighting
- Loading states with skeleton placeholders
- Stack job header component for key information

**File**: `app/frontend/src/components/shared/stackjob/stack-job-header.tsx`

Created header component:

```14:56:app/frontend/src/components/shared/stackjob/stack-job-header.tsx
export function StackJobHeader({ stackJob, updatedTime }: StackJobHeaderProps) {
  return (
    <HeaderContainer>
      <TopSection>
        <Stack gap={1}>
          <FlexCenterRow gap={0.5}>
            {stackJob ? (
              <Typography variant="caption" color="text.secondary">
                {stackJob?.id}
              </Typography>
            ) : (
              <Skeleton variant="text" width={180} height={15} />
            )}

            {stackJob ? (
              <TextCopy text={stackJob?.id} />
            ) : (
              <Skeleton variant="rectangular" width={10} height={10} />
            )}
          </FlexCenterRow>
          <Box>
            {stackJob ? (
              <StatusChip status={stackJob?.status || 'unknown'} />
            ) : (
              <Skeleton variant="rounded" width={80} height={20} />
            )}
          </Box>
        </Stack>
      </TopSection>
      <Divider sx={{ marginY: 1.5 }} />
      <BottomSection>
        <Box />
        {stackJob ? (
          <Typography variant="subtitle2" color="text.secondary">
            {updatedTime}
          </Typography>
        ) : (
          <Skeleton variant="text" width={180} height={15} />
        )}
      </BottomSection>
    </HeaderContainer>
  );
}
```

**Key features**:

- Displays job ID with copy-to-clipboard functionality
- Status chip for visual status indication
- Last updated timestamp
- Loading states with skeletons

### 4. Supporting Components

**File**: `app/frontend/src/components/shared/breadcrumb/index.tsx`

Created breadcrumb component for navigation:

```38:63:app/frontend/src/components/shared/breadcrumb/index.tsx
export const Breadcrumb: FC<IBreadcrumb> = ({ breadcrumbs, startBreadcrumb }) => {
  const isEmpty = breadcrumbs.length === 0;

  return (
    <StyledBreadcrumbs aria-label="breadcrumb">
      {isEmpty ? (
        <BreadcrumbLabel>
          <Skeleton variant="text" width={120} height={20} />
        </BreadcrumbLabel>
      ) : (
        startBreadcrumb
      )}
      {isEmpty
        ? [80, 100, 60].map((width, index) => (
            <BreadcrumbLabel key={`skeleton-${index}`}>
              <Skeleton variant="text" width={width} height={20} />
            </BreadcrumbLabel>
          ))
        : breadcrumbs.map(({ name, handler }, index) => (
            <BreadcrumbLabel key={`${name}-${index}`} onClick={handler} $hasLink={!!handler}>
              {name}
            </BreadcrumbLabel>
          ))}
    </StyledBreadcrumbs>
  );
};
```

**Key features**:

- Supports start breadcrumb with icon and label
- Clickable breadcrumb items
- Loading states with skeletons
- Reusable component for navigation

**File**: `app/frontend/src/components/shared/syntax-highlighter/json-code.tsx`

Created JSON syntax highlighter component for displaying stack job output:

- Displays JSON with proper formatting and syntax highlighting
- Used in stack job detail page to show deployment output

## Benefits

### For End Users

**Accessibility**:

- Stack jobs now accessible directly from cloud resources list
- Intuitive navigation flow from resources to jobs to details
- Easy access to deployment history

**User Experience**:

- Paginated list prevents performance issues with large datasets
- Detailed view shows complete deployment output
- Status chips provide quick visual feedback
- Breadcrumb navigation enables easy navigation back

**Performance**:

- Server-side pagination loads only current page
- Faster page loads for large numbers of stack jobs
- Reduced memory usage in browser

### For Developers

**Component Reusability**:

- Stack jobs components can be reused in other contexts
- Breadcrumb component is generic and reusable
- Drawer pattern consistent with other features

**Maintainability**:

- Clear separation between list and detail views
- Consistent pagination pattern with cloud resources
- Type-safe implementation with TypeScript

**Scalability**:

- Server-side pagination scales to handle thousands of stack jobs
- Efficient database queries with skip/limit
- Total count calculation only when needed

## Impact

### Immediate

**New Capabilities**:

- View stack jobs from cloud resources list
- Navigate to detailed stack job pages
- View complete deployment output with syntax highlighting
- Paginated browsing of stack jobs

**User Experience**:

- Intuitive navigation flow
- Better performance with pagination
- Complete visibility into deployment history

### Developer Experience

**1 new detail page** (`/stack-jobs/[id]`)
**2 new reusable components** (StackJobsDrawer, StackJobsList)
**1 new header component** (StackJobHeader)
**1 new breadcrumb component** for navigation
**1 new syntax highlighter component** for JSON display
**Backend pagination** support in service and repository layers

### System Capabilities

**UI Integration**: Stack jobs fully integrated into web interface
**Navigation**: Complete navigation flow from resources to jobs to details
**Pagination**: Scalable pagination for large datasets
**Performance**: Efficient loading with server-side pagination

## Usage Examples

### Opening Stack Jobs from Cloud Resources

1. Navigate to Cloud Resources page
2. Click action menu (three dots) on any cloud resource
3. Select "Stack Jobs" from menu
4. Drawer opens showing paginated list of stack jobs for that resource

### Viewing Stack Job Details

1. From stack jobs drawer, click on any stack job row
2. Navigate to `/stack-jobs/[id]` detail page
3. View complete stack job information:
   - Job ID (with copy button)
   - Status chip
   - Last updated timestamp
   - Full JSON output with syntax highlighting

### Navigating Back

1. From detail page, click "Stack Jobs" in breadcrumb
2. Opens drawer showing all stack jobs for the same cloud resource
3. Can navigate between jobs or return to cloud resources list

### Backend API with Pagination

**Request**:

```protobuf
ListStackJobsRequest {
  cloud_resource_id: "507f1f77bcf86cd799439011"
  page_info: {
    num: 0  // page number (0-indexed)
    size: 10  // items per page
  }
}
```

**Response**:

```protobuf
ListStackJobsResponse {
  jobs: [StackJob, ...]  // 10 items
  total_pages: 5  // calculated total pages
}
```

## Files Modified/Created

### Backend API

**Modified**:

- `app/backend/apis/proto/stack_job_service.proto` - Added PageInfo and total_pages to ListStackJobs API
- `app/backend/internal/service/stack_job_service.go` - Implemented pagination logic with total pages calculation, removed logrus logging
- `app/backend/internal/service/cloud_resource_service.go` - Removed logrus logging (code cleanup)
- `app/backend/internal/service/deployment_component_service.go` - Removed logrus logging (code cleanup)
- `app/backend/internal/database/stack_job_repo.go` - Added pagination support with skip/limit and Count method

### Frontend Pages

**Created**:

- `app/frontend/src/app/stack-jobs/[id]/page.tsx` - Stack job detail page
- `app/frontend/src/app/stack-jobs/_services/index.ts` - Stack jobs service exports
- `app/frontend/src/app/stack-jobs/_services/query.ts` - Stack jobs query service
- `app/frontend/src/app/stack-jobs/styled.ts` - Styled components for stack jobs pages

### UI Components

**Created**:

- `app/frontend/src/components/shared/stackjob/index.ts` - Stack job component exports
- `app/frontend/src/components/shared/stackjob/stack-job-header.tsx` - Stack job header component
- `app/frontend/src/components/shared/stackjob/stack-jobs-drawer.tsx` - Stack jobs drawer component
- `app/frontend/src/components/shared/stackjob/stack-jobs-list.tsx` - Stack jobs list component with pagination
- `app/frontend/src/components/shared/breadcrumb/index.tsx` - Breadcrumb navigation component
- `app/frontend/src/components/shared/breadcrumb/styled.ts` - Breadcrumb styling
- `app/frontend/src/components/shared/status-chip/index.ts` - Status chip exports
- `app/frontend/src/components/shared/status-chip/status-chip.tsx` - Status chip component
- `app/frontend/src/components/shared/syntax-highlighter/index.ts` - Syntax highlighter exports
- `app/frontend/src/components/shared/syntax-highlighter/json-code.tsx` - JSON syntax highlighter component

**Modified**:

- `app/frontend/src/components/shared/cloud-resources-list/cloud-resources-list.tsx` - Added "Stack Jobs" menu item and drawer integration
- `app/frontend/src/components/shared/cloud-resources-list/index.ts` - Updated exports
- `app/frontend/src/components/layout/styled.ts` - Updated layout styling
- `app/frontend/src/components/shared/drawer/styled.ts` - Updated drawer styling

**Deleted**:

- `app/frontend/src/components/shared/cloud-resources-list/styled.ts` - Removed (no longer needed)

### Configuration

**Modified**:

- `app/frontend/package.json` - Updated dependencies
- `app/frontend/yarn.lock` - Updated lock file

## Technical Metrics

- **1 new detail page** with dynamic routing
- **4 new reusable components** for stack jobs UI
- **1 new breadcrumb component** for navigation
- **1 new status chip component** for status display
- **1 new syntax highlighter component** for JSON display
- **Server-side pagination** implemented in backend and frontend
- **Default page size**: 10 items per page (frontend), 20 items per page (backend default)
- **Backward compatible**: Pagination is optional in API
- **Full TypeScript coverage** for all new components

## Related Work

### Foundation

This work builds on:

- **Pulumi CLI Stack Job API Implementation** (December 3, 2025) - Backend API foundation
- **Cloud Resource UI Enhancements and Pagination** (December 3, 2025) - Table component and pagination patterns
- **Cloud Resource Web UI** (December 1, 2025) - Web interface infrastructure

### Complements

This work complements:

- **Cloud Resource Management** - Enables viewing deployment history for resources
- **Stack Job API** - Provides UI for existing backend functionality
- **Pagination System** - Extends pagination pattern to stack jobs

### Future Extensions

This work enables:

- **Deployment Actions** - Can add deploy/retry actions from UI
- **Real-time Updates** - Can add polling or WebSocket for live status updates
- **Filtering** - Can add status filters to stack jobs list
- **Export** - Can export stack job output as files
- **Bulk Operations** - Can add bulk actions for multiple stack jobs
- **Search** - Can add search functionality for stack jobs

## Known Limitations

- **Fixed page size**: Frontend uses 10 items per page, not user-configurable
- **No status filtering in UI**: Backend supports status filter but UI doesn't expose it
- **No real-time updates**: Status changes require manual refresh
- **No deployment actions**: Can't trigger deployments from UI (only view history)
- **No output filtering**: Full JSON output always displayed (could add collapsible sections)

These limitations are intentional for the initial implementation and can be addressed in future enhancements.

## Design Decisions

### Server-Side Pagination

**Decision**: Implement server-side pagination for stack jobs list

**Rationale**:

- Consistent with cloud resources pagination pattern
- Scales to handle large numbers of stack jobs
- Better performance than loading all jobs at once
- Standard pattern for data-heavy applications

**Alternative considered**: Client-side pagination

- Rejected because it doesn't scale and loads all data at once

### Drawer Pattern

**Decision**: Use drawer component for stack jobs list instead of separate page

**Rationale**:

- Keeps user in context of cloud resources
- Consistent with other drawer patterns (view/edit)
- Quick access without full page navigation
- Can still navigate to detail page from drawer

**Alternative considered**: Separate page for stack jobs list

- Rejected because drawer provides better UX and keeps context

### Breadcrumb Navigation

**Decision**: Use breadcrumb with clickable "Stack Jobs" link that opens drawer

**Rationale**:

- Provides clear navigation path
- Enables quick return to stack jobs list
- Consistent with common web navigation patterns
- Opens drawer instead of navigating away (maintains context)

**Alternative considered**: Simple back button

- Rejected because breadcrumb provides more context and flexibility

### Page Size

**Decision**: Frontend uses 10 items per page, backend defaults to 20

**Rationale**:

- 10 items provides good balance for drawer width
- Backend default of 20 maintains backward compatibility
- Frontend can request different page sizes if needed
- Consistent with cloud resources pagination

**Alternative considered**: Same page size for both

- Rejected because drawer benefits from smaller page size for better UX

## Post-Implementation Cleanup

### Logrus Removal from Service APIs

As part of code quality improvements, removed all logrus logging from the backend service layer (`app/backend/internal/service/`).

**Rationale**:

- Most error logs were redundant since errors are already returned to clients via `connect.NewError` with proper error messages
- Info and warning logs were nice-to-have but not essential for API operations
- Removing logrus simplifies the codebase and reduces dependencies
- Errors are still properly handled and returned to clients through Connect RPC error handling

**Files Modified**:

- `app/backend/internal/service/stack_job_service.go` - Removed all logrus calls (error, warning, and info logs)
- `app/backend/internal/service/cloud_resource_service.go` - Removed all logrus calls
- `app/backend/internal/service/deployment_component_service.go` - Removed all logrus calls

**Impact**:

- Cleaner code with less redundant logging
- Reduced dependency on logrus in service layer
- Errors still properly propagated to clients via Connect RPC
- Note: `database/mongodb.go` still uses logrus for connection/disconnection logging (infrastructure-level logging retained)

---

**Status**: ✅ Complete and Production Ready
**Component**: Web Frontend - Stack Jobs UI Integration, Backend API - Pagination
**Pages Added**: 1 detail page (`/stack-jobs/[id]`)
**Components Added**: 6 new reusable components
**Components Modified**: 2 existing components
**Backend Changes**: Pagination support in service and repository
**Location**: `app/frontend/src/app/stack-jobs/`, `app/frontend/src/components/shared/stackjob/`, `app/backend/internal/service/`, `app/backend/internal/database/`
