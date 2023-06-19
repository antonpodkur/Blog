import { SubmitHandler, useForm } from "react-hook-form"
import { useAxios } from "../api/axios"
import { useNavigate } from "react-router-dom"
import { useEffect, useState } from "react"

interface IFormInput {
  name: string,
  email: string,
  password: string,
  passwordConfirm: string
}

const Register: React.FC = () => {
  const { register, watch, formState: { errors }, handleSubmit } = useForm<IFormInput>()
  const [errMsg, setErrMsg] = useState("")
  const axios = useAxios()
  const navigate = useNavigate()

  const watchName = watch("name")
  const watchEmail = watch("email")
  const watchPassword = watch("password")
  const watchPasswordConf = watch("passwordConfirm")

  useEffect(() => {
    console.log(errMsg)
    setErrMsg("")
  }, [watchName, watchEmail, watchPassword, watchPasswordConf])

  const onSubmit: SubmitHandler<IFormInput> = async (data) => {
    try {
      const response = await axios.post("/api/v1/auth/register", data)
      if (response.status !== 201) {
        setErrMsg(response.data.message)
      }
      navigate('/login')
    } catch (error) {
      // if (error instanceof AxiosError) {
      //   console.log(error.response?.data.message)
      //   setErrMsg(error.response?.data.message ?? "")
      // }
      // else {
      //   setErrMsg("Ooops. Something went wrong")
      // }

      //TODO: Add min and max width validation
      console.log("Sign Up failed")
    }
  }

  return (
    <div className="flex flex-col items-center justify-center">
      <div className="my-4 md:my-12 text-center text-4xl md:text-6xl font-extrabold">Lets create your account</div>
      {errMsg !== "" && <h1 className="text-lg">{errMsg}</h1>}
      <form className="form-control w-full max-w-xs md:max-w-md mt-6" onSubmit={handleSubmit(onSubmit)}>
        <label className="label label-textl font-bold">Name</label>
        <input type="text" className="input input-bordered w-full max-w-xs md:max-w-md" placeholder="Enter your name" {...register("name", { required: true, maxLength: 255 })} />
        {errors.name && errors.name.type === "required" && <label className="label label-text-alt text-red">Name is required</label>}

        <label className="label label-textl font-bold">Email</label>
        <input type="email" className="input input-bordered w-full max-w-xs md:max-w-md" placeholder="Enter your email" {...register("email", { required: true, maxLength: 255 })} />
        {errors.email && errors.email.type === "required" && <label className="label label-text-alt text-red">Email is required</label>}

        <label className="label label-textl font-bold">Password</label>
        <input type="password" className="input input-bordered w-full max-w-xs md:max-w-md" placeholder="Enter your password" {...register("password", { required: true, maxLength: 255 })} />
        {errors.password && errors.password.type === "required" && <label className="label label-text-alt text-red">Password is required</label>}

        <label className="label label-textl font-bold">Password Confirmation</label>
        <input type="password" className="input input-bordered w-full max-w-xs md:max-w-md" placeholder="Repeat your password" {...register("passwordConfirm", { required: true, maxLength: 255 })} />
        {errors.passwordConfirm && errors.passwordConfirm.type === "required" && <label className="label label-text-alt text-red">Password confirmation is required</label>}

        <input className="btn btn-primary font-bold mt-6" type="submit" />
      </form>
    </div >
  )
}

export default Register
