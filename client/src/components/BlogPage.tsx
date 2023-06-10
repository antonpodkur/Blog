import { useEffect, useState } from "react"
import { Article } from "../models/article"
import parse from "html-react-parser"
import axios from "axios"
import { useParams } from "react-router-dom"

const BlogPage: React.FC = () => {
    const {id} = useParams()
    const [article, setArticle] = useState<Article | null>(null)
    const [isLoading, setIsLoading] = useState(true)

    useEffect(() => {
        async function fetchArticle(): Promise<Article> {
            const result = await axios.get(`http://localhost:4000/api/v1/article/${id}`);
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
            <div>{parse(article!.content)}</div>
        </div>    
    )
}

export default BlogPage;
