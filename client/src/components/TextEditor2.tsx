import { useEffect, useState } from "react"
import { Editor } from "react-draft-wysiwyg"
import { EditorState } from "draft-js"
import "react-draft-wysiwyg/dist/react-draft-wysiwyg.css"

const TextEditor2: React.FC = () => {
    const [editorState, setEditorState] = useState(() => EditorState.createEmpty())

    return (
        <div>
            <Editor
              editorState={editorState}
              toolbarClassName="toolbarClassName"
              wrapperClassName="wrapperClassName"
              editorClassName="editorClassName"
              onEditorStateChange={setEditorState}
            />
        </div>
    )
}

export default TextEditor2
