import axios from "axios";
import React, { useEffect, useState } from "react";
import './App.css';


const App = () => {
  const [isOpen, setIsOpen] = useState(false)
  const [isOpenGroup, setIsOpenGroup] = useState(false)
  const [title, setTitle] = useState("")
  // const [completed, setCompleted] = useState(false)
  const [todos, setTodos] = useState([]);

  useEffect(() => {
    getTodos();
  }, []);

  const getTodos = async () => {
    try {
      const res = await axios.get("http://localhost:5000/check-todo");
      setTodos(res.data);
    } catch (error) {
      console.log(error);
    }
  }

  const createTodo = async (e) => {
    e.preventDefault()
    try {
      const response = await axios.post("http://localhost:5000/create-todo", {title: title})
      setTitle("")
      getTodos()
      console.log(response.data)
    } catch (error) {
      console.log(error)
    }
  }

  const handleDelete = async (id) => {
    try {
      const res = await axios.delete(`http://localhost:5000/delete-todo/${id}`)
      getTodos()
      console.log(res.data)
    } catch (error) {
      console.log(error)
    }
  }

  const handleCompleted = async (id, title, completed) => {
    try {
      const res = await axios.put(`http://localhost:5000/update-todo/${id}`, { id, title, completed });
      getTodos();
      console.log(res.data)
    } catch (error) {
      console.log(error);
    }
  }

  const handleOpenModal = () => {
    setIsOpen(true)
  }
  
  const handleCloseModal = () => {
    setIsOpen(false)
  }

  const handleOpenModalGroup = () => {
    setIsOpenGroup(true)
  }

  const handleCloseModalGroup = () => {
    setIsOpenGroup(false)
  }

  return (
    <div className="wrapper">
      <header className="header">
        <h1 className="header-logo">todo</h1>
        <button onClick={handleOpenModal} className="addBtn add-todo">+</button>
      </header>
      <div className="container">
        <div className="groups">
          <div className="add-group" onClick={handleOpenModalGroup}>
            <button className="addBtn add-group-btn">+</button>
            <div className="add-group-text">Добавить группу</div>
          </div>
          <div className="group-block">
            
          </div>
          <div className="isDone">
            <input className="hide-checkbox" name="hide-checkbox" type="checkbox" />
            <p>Hide Done Tasks</p>
          </div>
        </div>
        {isOpenGroup && (
          <div className="modal-overlay">
            <div className="modal">
              <div className="modal-header">
                <h2>Добавить Группу</h2>
                <button onClick={handleCloseModalGroup} className="close-modal">&times;</button>
              </div>
              <div className="modal-body">
                <p>Введите имя группы: </p>
              </div>
              <form className="modal-footer">
                <button type="submit" className="btn primary">Добавить</button>
              </form>
            </div>
          </div>
        )}
        <div className="todos-container">
          {todos && todos.map(todo => (
            <div className="todos" key={todo.ID}>  
              <h2 className="title">{todo.Title}</h2>
              <div className="todoIsDone">
                {todo.Completed.toString()}
                <input 
                  checked={todo.Completed} 
                  onChange={() => handleCompleted(todo.ID, todo.Title, !todo.Completed)} 
                  type="checkbox"
                  name="checkbox"
                />
              </div>
              <button onClick={() => handleDelete(todo.ID)}>Delete</button>
            </div>
          ))}
        </div>
        {/* <div className="todos-container"></div> */}
      </div>
      {isOpen && (
        <div className="modal-overlay">
          <div className="modal">
            <div className="modal-header">
              <h2>Добавить заметку</h2>
              <button onClick={handleCloseModal} className="close-modal">&times;</button>
            </div>
            <div className="modal-body">
              <input type="text" value={title} onChange={e => setTitle(e.target.value)} />
              {/* <p>Modal content goes here.</p> */}
            </div>
            <form className="modal-footer" onSubmit={createTodo}>
              <button type="submit" className="btn primary">Добавить</button>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

export default App;
