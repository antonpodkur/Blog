export interface Article {
    id: string | null,
    title: string,
    content: string,
    createdAt: Date | null, 
    updatedAt: Date | null, 
    userId: string
}
