import { useEffect, useState } from "react"
import { Article } from "../models/article"
import axios from "axios"
import { Link } from "react-router-dom"
import { useAuthStore } from "../store/store"

const AllArticles: React.FC = () => {
  const [articles, setArticles] = useState<Article[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const isLoggedIn = useAuthStore(store => store.isLoggedIn)
  const user = useAuthStore(store => store.user)

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
      console.log(user, isLoggedIn)

      await fetchArticles()
    })()
  }, [])

  if (isLoading) {
    return <div>Loading...</div>
  }

  return (
    <div className="flex flex-col w-full items-center justify-center">
      <div className="p-4 m-4 font-bold text-4xl md:text-6xl">Articles</div>
      {articles.length > 0 &&
        <ul className="flex flex-col w-3/4 md:min-w-3/12 md:w-3/6">
          {articles.map((article) => (
            <li key={article.id} className="w-full flex flex-col">
              <Link to={`/article/${article.id}`} className=" py-5 px-2 m-3 border-gray-950 ease-in-out duration-300 border-2 hover:border-4  font-semibold">
                {article.title}
              </Link>
            </li>
          ))}
        </ul>
      }
    </div>
  )
}

export default AllArticles
