import axios from "axios";
import React, { useState } from "react";
import './App.css';


const App = () => {
  const [user, setUser] = useState("")
  const [name, setName] = useState("")
  const [message, setMessage] = useState("")

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      const res = await axios.get(`http://localhost:5000/${name}`)
      setUser(res.data)
      console.log(res)
      console.log(res.data.name)
    } catch (error) {
      console.log(error)
      setMessage("An error occured")
    }
  }

  // useEffect(() => {
  //   axios.get("http://localhost:5000")
  //   .then(res => {
  //     setUser(res.data)
  //   })
  //   .catch(err => console.log(err))
  // }, [])

  return (
    <div className="App">
      <div>
        <form onSubmit={handleSubmit}>
          <label>
            Enter your name:
            <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
          </label>
          <button type="submit">Submit</button>
        </form>
        {message && <p>{message}</p>}

        <div>id: {user.id}</div>
        <div>name: {user.name}</div>
      </div>
    </div>
  );
}

export default App;
