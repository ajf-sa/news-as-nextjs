// "id": 1,
// "name": "محليات",
// "slug": "local",
// "published_at": "2021-05-15T15:09:32.221Z",
// "created_at": "2021-05-15T15:09:29.637Z",
// "updated_at": "2021-05-15T15:25:51.498Z"

export interface Tag {
  id: string | null;
  name: string | null;
  slug: string | null;
}

export interface PostType {
  id: string | null;
  title: string | null;
  slug: string | null;
  description: string | null;
  created_at: string | null;
  image: ImageType | null;
}
export interface ImageType {
  url: string | null;
}
