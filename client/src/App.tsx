import { Route, Routes } from 'react-router-dom'
import BlogPage from './components/BlogPage'
import TextEditor from './components/TextEditor'
import AllArticles from './pages/AllArticles'

function App() {
    return (
        <div>
            <Routes>
                <Route path="/" element={<AllArticles/>} />
                <Route path="/create" element={<TextEditor/>} />
                <Route path="/article/:id" element={<BlogPage/>} />
            </Routes>
        </div> 
    )
}

export default App
