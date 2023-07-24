import { Route, Routes } from 'react-router-dom'
import BlogPage from './components/BlogPage'
import TextEditor from './components/TextEditor'
import AllArticles from './pages/AllArticles'
import Register from './pages/Register'
import Login from './pages/Login'
import { useAuthStore } from './store/store'
import Protected from './components/Protected'
import Navbar from './components/Navbar'

function App() {
  const isLoggedIn = useAuthStore(store => store.isLoggedIn)

  return (
    <div>
      <Navbar/>
      <Routes>
        <Route path="/" element={<AllArticles />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/create" element={
          <Protected isLoggedIn={isLoggedIn} >
            <TextEditor />
          </Protected>
        } />
        <Route path="/article/:id" element={<BlogPage />} />
      </Routes>
    </div>
  )
}

export default App
