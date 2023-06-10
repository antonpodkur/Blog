import axios from "axios"
import { useEffect, useState } from "react"
import ReactQuill from "react-quill"
import 'react-quill/dist/quill.snow.css'
import { Article } from "../models/article"

const TextEditor: React.FC = () => {
    const [value, setValue] = useState<string>('')

    const CreateArticle = async () => {
        const article: Article = {
            id: null,
            content: value,
            createdAt: null,
            updatedAt: null,
            }
        const result = await axios.post("http://localhost:4000/api/v1/article/", article);   
        if (result.status !== 201) {
            console.log("Error creating article")
        }
        console.log(result.data.data)
    }

    useEffect(() => {
        console.log(value)

    }, [value])
    
    return (
        <div>
            <ReactQuill theme="snow" value={value} onChange={setValue} />
            <button onClick={async () => CreateArticle()}>Create</button>
        </div>
    )
}

export default TextEditor
