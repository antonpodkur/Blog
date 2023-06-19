import SimpleMDEditor from 'react-simplemde-editor'
import 'easymde/dist/easymde.min.css'
import { useMemo, useState } from 'react'
import { Article } from '../models/article'
import axios from 'axios'
import { useAuthStore } from '../store/store'
import { User } from '../models/user'
import EasyMDE from 'easymde'

type ImageUploadType = {
  (image: File,
    onSuccess: (url: string) => void,
    onError: (errorMessage: string) => void): void
}

const TextEditor: React.FC = () => {
  const [title, setTitle] = useState("")
  const [value, setValue] = useState("")
  const user: User | null = useAuthStore((state) => state.user)

  const handleTitleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setTitle(e.target.value)
  }

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
      const imageUrl = res.data.data.file
      onSuccess(imageUrl)
    } catch (error) {
      return onError(error as string)
    }
  }

  const CreateArticle = async () => {
    const article: Article = {
      id: null,
      title: title,
      content: value,
      createdAt: null,
      updatedAt: null,
      userId: user?.id ?? ""
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
      <input placeholder='Enter your post title' value={title} onChange={handleTitleChange} />
      <SimpleMDEditor
        id='editor'
        value={value}
        onChange={handleValueChange}
        options={newOptions as EasyMDE.Options}
      />
      <button onClick={() => CreateArticle()}>Save</button>
    </div>
  )
}

export default TextEditor
