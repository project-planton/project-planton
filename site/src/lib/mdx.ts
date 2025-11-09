export interface Author {
  name: string;
  role?: string;
  image?: string;
  url?: string;
}

export class MDXParser {
  /**
   * Reconstructs MDX content from markdown with frontmatter
   * This is a simple pass-through for now since we're using react-markdown
   */
  static reconstructMDX(content: string): string {
    return content;
  }
}

