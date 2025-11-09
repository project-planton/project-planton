import path from 'path';

const CONTENT_DIRECTORIES = {
  DOCS: path.join(process.cwd(), 'public/docs'),
} as const;

export const DOCS_DIRECTORY = CONTENT_DIRECTORIES.DOCS;

