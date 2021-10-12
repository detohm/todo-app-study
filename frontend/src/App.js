import { Button } from "./components/atoms/button";
import { Card } from "./components/atoms/card";
import { Title } from "./components/atoms/title";
import "./style.scss";
function App() {
  return (
    <div className="app">
      <Title title="To-Do List App" />
      <Button label="click!" />
      <Card>
        <Button label="card button" />
      </Card>

    </div>
  );
}

export default App;
