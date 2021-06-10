export interface Tag {
  id: string;
  name: string;
  slug: string;
}

export interface PostType {
  id: string;
  slug: string;
  uuid: string;
  title: string;
  url: string;
  html: string;
  excerpt: string;
  comment_id: string;
  feature_image: string;
  created_at: string;
}
export interface ImageType {
  url: string;
}
