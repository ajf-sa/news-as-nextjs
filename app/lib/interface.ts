// "id": 1,
// "name": "محليات",
// "slug": "local",
// "published_at": "2021-05-15T15:09:32.221Z",
// "created_at": "2021-05-15T15:09:29.637Z",
// "updated_at": "2021-05-15T15:25:51.498Z"

export interface Tag{
    id : number | null
    name : string | null
    slug : string | null
}

export interface PostType{
    id: number | null
    title : string | null
    description :string | null
    tags : Array<Tag> 
  
}