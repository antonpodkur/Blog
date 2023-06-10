import { useEffect, useState } from "react"
import { Article } from "../models/article"
import axios from "axios"
import { Link } from "react-router-dom"

const AllArticles: React.FC = () => {
    const [articles, setArticles] = useState<Article[]>([])
    const [isLoading, setIsLoading] = useState(true)

    const fetchArticles = async () => {
        const result = await axios.get("http://localhost:4000/api/v1/article/")
        if (result.status !== 200) {
            console.log("Failed to fetch articles")
        }
        setArticles(result.data.data)
        setIsLoading(false)
    }

    useEffect(() => {
        (async () => {
            await fetchArticles()
        })()
    }, [])

    if (isLoading) {
        return <div>Loading...</div>         
    }

    return (
        <div>
            {articles.length > 0 &&
                <ul>
                    {articles.map((article) => (
                        <li key={article.id}>
                            <Link to={`/article/${article.id}`}>
                                {article.id}
                            </Link>
                        </li>
                    ))}
                </ul>
            }
        </div>
    )
}

export default AllArticles
