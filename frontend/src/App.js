import { Button } from "./components/atoms/button";
import { Title } from "./components/atoms/title";
import "./style.scss";
function App() {
  return (
    <div className="app">
      <Title title="To-Do List App" />
      <Button label="click!" />
    </div>
  );
}

export default App;
