import { Button } from "./components/atoms/button";
import { Card } from "./components/atoms/card";
import { Header } from "./components/organisms/header";
import "./style.scss";
function App() {
  return (
    <div className="app">
      <Header />
      <Button label="click!" />
      <Card>
        <Button label="card button" />
      </Card>

    </div>
  );
}

export default App;
