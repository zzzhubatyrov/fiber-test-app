import axios from "axios";
import React, { useEffect, useState } from "react";

function TodoItem({ todo }) {
  const [isEditing, setIsEditing] = useState(false);
  const [title, setTitle] = useState(todo.Title);
  const [description, setDescription] = useState(todo.Description);

  useEffect(() => {
    setTitle(todo.Title);
    setDescription(todo.Description);
  }, [todo]);

  const handleEditClick = () => {
    setIsEditing(true);
  };

  const handleSaveClick = async () => {
    try {
      const updatedTodo = {
        ...todo,
        Title: title,
        Description: description,
      };
      const res = await axios.put(`http://localhost:5000/update-todo/${todo.ID}`, updatedTodo);
      console.log(res.data);
      setIsEditing(false);
    } catch (error) {
      console.log(error);
    }
  };

  const handleCompleted = async () => {
    try {
      const updatedTodo = {
        ...todo,
        Completed: !todo.Completed,
      };
      const res = await axios.put(`http://localhost:5000/update-todo/${todo.ID}`, updatedTodo);
      console.log(res.data);
    } catch (error) {
      console.log(error);
    }
  };

  const handleDelete = async () => {
    try {
      const res = await axios.delete(`http://localhost:5000/delete-todo/${todo.ID}`);
      console.log(res.data);
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <div className="todos" key={todo.ID}>
      {isEditing ? (
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
      ) : (
        <h2 className="title">{title}</h2>
      )}

      <div className="descriptionBlock">
        {isEditing ? (
          <input
            type="text"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          />
        ) : (
          <h3 className="description">{description}</h3>
        )}
      </div>

      <div className="editTask">
        {isEditing ? (
          <button onClick={handleSaveClick}>Save</button>
        ) : (
          <button onClick={handleEditClick}>Edit Task</button>
        )}
      </div>

      <div className="todoIsDone">
        {todo.Completed.toString()}
        <input
          checked={todo.Completed}
          onChange={handleCompleted}
          type="checkbox"
          name="checkbox"
        />
      </div>

      <button onClick={handleDelete}>Delete</button>
    </div>
  );
}

export default TodoItem;
 