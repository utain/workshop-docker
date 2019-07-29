import React from "react";
import "./App.css";

function App() {
  let [todos, setTodos] = React.useState([]);
  React.useEffect(() => {
    (async () => {
      let res = await fetch("/api/tasks");
      let tasks = await res.json();
      setTodos(tasks);
    })();
  }, []);
  return (
    <div className="App">
      <header className="App-header">
        <p>Todo App</p>
        <ul>
          {todos.map(item => {
            return (
              <li key={item.title} className={item.completed ? "done" : ""}>
                {item.title}
              </li>
            );
          })}
        </ul>
      </header>
    </div>
  );
}

export default App;
