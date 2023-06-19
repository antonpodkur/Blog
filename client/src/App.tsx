import { Route, Routes } from 'react-router-dom'
import BlogPage from './components/BlogPage'
import TextEditor from './components/TextEditor'
import AllArticles from './pages/AllArticles'
import Register from './pages/Register'
import Login from './pages/Login'

function App() {
    return (
        <div>
            <Routes>
                <Route path="/" element={<AllArticles/>} />
                <Route path="/register" element={<Register/>} />
                <Route path="/login" element={<Login/>} />
                <Route path="/create" element={<TextEditor/>} />
                <Route path="/article/:id" element={<BlogPage/>} />
            </Routes>
        </div> 
    )
}

export default App
