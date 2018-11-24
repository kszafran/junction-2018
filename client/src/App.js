import React, { Component } from "react";
import Header from "./components/Header";
import DeviceList from "./components/DeviceList";
import DeviceDetails from "./components/DeviceDetails";
import { Route } from "react-router-dom";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { devices: [] };
  }

  componentDidMount = async () => {
    const MOCK_DATA_URL =
      "https://gist.githubusercontent.com/vietdang7/a5d495df7bc398a1f6c46eade07d452d/raw/d4f7639672d99e1556300c4404356d436e64e854/Junction%2520Data";

    const result = await fetch(MOCK_DATA_URL);
    const devices = await result.json();

    this.setState({
      devices
    });
  };

  getDeviceFromID = id => {
    return this.state.devices.find(e => e.id === id);
  };

  render() {
    const { devices } = this.state;

    return (
      <div className="App">
        <Header />
        <div className="wrapper">
          <Route exact path="/" render={() => <DeviceList devices={devices} />} />
          <Route
            path="/device/:id"
            render={props => (
              <DeviceDetails
                {...props}
                getDeviceFromID={this.getDeviceFromID}
              />
            )}
          />
        </div>
      </div>
    );
  }
}

export default App;
