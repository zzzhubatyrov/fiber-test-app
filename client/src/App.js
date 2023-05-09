import axios from "axios";
import React, { useEffect, useState } from "react";
import './App.css';


const App = () => {

  const [value, setValue] = useState("")

  // const handleSubmit = async (e) => {
  //   e.preventDefault()
  //   try {
  //     const res = await axios.get(`http://localhost:5000/${name}`)
  //     setUser(res.data)
  //     console.log(res)
  //   } catch (error) {
  //     console.log(error)
  //     setMessage("An error occured")
  //   }
  // }


  // create-todo
  // create-models
  // check-todo
  // "/" - main page

  useEffect(() => {
    axios.get("http://localhost:5000")
    .then(res => {
      setUser(res.data)
    })
    .catch(err => console.log(err))
  }, [])

  return (
    <div className="App">
      <div>
        <p>ToDo's</p>
        <div key={id}>

        </div>
      </div>
    </div>
  );
}

export default App;
