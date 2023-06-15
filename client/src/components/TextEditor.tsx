import  SimpleMDEditor  from 'react-simplemde-editor'
import 'easymde/dist/easymde.min.css'
import { useMemo, useState } from 'react'
import { Article } from '../models/article'
import axios from 'axios'

type ImageUploadType = {
   (image: File, 
   onSuccess: (url: string) => void, 
   onError: (errorMessage: string) => void) : void
}

const TextEditor: React.FC = () => {
    const [value, setValue] = useState("")

    const handleValueChange = (value: string) => {
        console.log(value)
        setValue(value)
    }

    const imageUpload: ImageUploadType = async (image, onSuccess, onError) => {
        try {
            const data = new FormData()
            data.append("file", image)
            const res = await axios.post("http://localhost:4000/api/v1/files/", data)
            if (res.status !== 201) {
                    throw new Error(res.data.message)
            }
            const imageUrl = `http://localhost:4000/api/v1/files/${res.data.data}`
            console.log(imageUrl)
            onSuccess(imageUrl)
        } catch (error) {
            return onError(error)
        } 
    }

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

    const newOptions = useMemo(() => {
        return {
            spellChecker: false,
            showIcons: ["strikethrough", "table", "code", "upload-image"],
            hideIcons: ["image"],
            uploadImage: true,
            imageUploadFunction: imageUpload
        }
    }, [])

    return (
        <div>
            <SimpleMDEditor
                id='editor'
                value={value}
                onChange={handleValueChange}
                options = {newOptions}
            />
            <button onClick={() => CreateArticle()}>Save</button>
        </div>
    ) 
}

export default TextEditor
