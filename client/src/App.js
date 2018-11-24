import React, { Component } from "react";
import Header from "./components/Header";
import DeviceList from "./components/DeviceList";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { sideNav: false };
  }

  render() {
    return (
      <div className="App">
        <Header />
        <div className="wrapper">
          <DeviceList />
        </div>
      </div>
    );
  }
}

export default App;
