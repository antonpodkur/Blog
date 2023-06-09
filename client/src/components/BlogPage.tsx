import { useEffect, useState } from "react"
import { Article } from "../models/article"
import parse from "html-react-parser"
import axios from "axios"

const BlogPage: React.FC = () => {
    const [article, setArticle] = useState<Article | null>(null)
    const [isLoading, setIsLoading] = useState(true)

    useEffect(() => {
        async function fetchArticle(): Promise<Article> {
            const result = await axios.get("http://localhost:4000/api/v1/article/post");
            if (result.status !== 200 && result.data.status !== "success") {
                throw new Error("Error while fetching blogpost")
            }
            return result.data.data as Article
        }

        (async () => {

            console.log("I am here")
            try {
                const article = await fetchArticle() 
                console.log(article)
                setArticle(article)
                setIsLoading(false)
            } catch (e) {
                console.log(e)
            }
        })();

    console.log("Hello from page") 
    }, [])

    if (isLoading) {
        return (
            <h1>Loading</h1>
        )
    }

    return (
        <div className="flex flex-col items-center justify-center ">
            <h1 className="text-2xl">Blog Page</h1>
            <div>{parse(article!.content)}</div>
        </div>    
    )
}

export default BlogPage;
