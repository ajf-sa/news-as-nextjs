// "id": 1,
// "name": "محليات",
// "slug": "local",
// "published_at": "2021-05-15T15:09:32.221Z",
// "created_at": "2021-05-15T15:09:29.637Z",
// "updated_at": "2021-05-15T15:25:51.498Z"

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
