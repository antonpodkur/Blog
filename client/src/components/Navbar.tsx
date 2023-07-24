import { Link, useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/store";
import { useAxios } from "../api/axios";

const Navbar: React.FC = () => {
  const isLoggedIn = useAuthStore(store => store.isLoggedIn)
  const user = useAuthStore(store => store.user)
  const reset = useAuthStore(store => store.reset)
  const axios = useAxios()
  const navigate = useNavigate()

  const handleLogOut = async () => {
    try {
      const res = await axios.get("/api/v1/auth/logout");
      reset()
      navigate("/login")
    }
    catch (error) {
      console.log(error)
    }
  }

  return (
    <div className="navbar bg-base-100">
      <div className="navbar-start">
        <div className="dropdown">
          <label tabIndex={0} className="btn btn-ghost lg:hidden">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-5 w-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h8m-8 6h16"
              />
            </svg>
          </label>
          <ul
            tabIndex={0}
            className="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
          >
            <li>
              <Link to="/create">Create Post</Link>
            </li>
            <li>
              <a>Item 2</a>
            </li>
          </ul>
        </div>
        <Link to="/" className="btn btn-ghost normal-case text-xl">BlogMaker</Link>
      </div>
      <div className="navbar-center hidden lg:flex">
        <ul className="menu menu-horizontal px-1">
          <li>
              <Link to="/create">Create Post</Link>
          </li>
          <li>
            <a>Item 2</a>
          </li>
        </ul>
      </div>
      <div className="navbar-end mr-2">
        { isLoggedIn && 
          <ul className="flex items-center">
            <li>
              <div className="font-bold text-lg">{user!.name}</div>
            </li>
            <li className="ml-4">
              <div className="btn btn-sm" onClick={ async () => await handleLogOut()}>Log out</div>
            </li>
          </ul>
        }
        { !isLoggedIn && 
          <ul className="flex items-center">
            <li>
              <Link to="/register" className="btn btn-sm">Sign Up</Link>
            </li>
            <li className="ml-4">
              <Link to="/login" className="btn btn-sm">Sign In</Link>
            </li>
          </ul>
        }
      </div>
    </div>
  );
};

export default Navbar;
