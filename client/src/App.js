import axios from "axios";
import React, { useEffect, useState } from "react";
import './App.css';
import './components/todo/todo.js';
import TodoItem from "./components/todo/todo.js";

const App = () => {
  const [isOpen, setIsOpen] = useState(false)
  const [isOpenGroup, setIsOpenGroup] = useState(false)
  const [title, setTitle] = useState("")
  const [description, setDescription] = useState("")
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
      const response = await axios.post("http://localhost:5000/create-todo", {title: title, description: description})
      setTitle("")
      setDescription("")
      getTodos()
      console.log(response.data)
    } catch (error) {
      console.log(error)
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
                <button disabled={true} type="submit" className="btn primary">Добавить</button>
              </form>
            </div>
          </div>
        )}
        <div className="todos-container">
          {todos && todos.map(todo => (
            <div key={todo.ID}>
              <TodoItem todo={todo} />
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
              <input placeholder="Title" type="text" value={title} onChange={e => setTitle(e.target.value)} />
              <input placeholder="Description" type="text" value={description} onChange={e => setDescription(e.target.value)} />
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
