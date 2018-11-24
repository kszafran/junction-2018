import React, { Component } from "react";
import Header from "./components/Header";
import DeviceList from "./components/DeviceList";
import DeviceDetails from "./components/DeviceDetails";
import { Route } from "react-router-dom";
import "./App.css";
import Loader from "./components/Loader";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = { devices: [], loader: true };
  }

  fetchJSON = async url => {
    const response = await fetch(url);
    return await response.json();
  };

  componentDidMount = async () => {
    // return only 1 object from the API
    const DNA_C_DATA =
      "https://cors-anywhere.herokuapp.com/https://peaceful-shelf-99858.herokuapp.com/client-health/3c:97:0e:da:3c:65";

    const MOCK_DATA_URL =
      "https://gist.githubusercontent.com/ashleynguci/5fd6c84358844f9ac50f713b039bad8f/raw/a1cde2ca3787f31cb960c591b27f70d6634f2145/mock.json";
    const urlsArray = [MOCK_DATA_URL, DNA_C_DATA];
    const promiseArray = urlsArray.map(url => this.fetchJSON(url));
    const dataArray = await Promise.all(promiseArray);

    const devices = [...dataArray[0], dataArray[1]];

    this.setState({
      devices: devices,
      loader: false
    });
  };

  getDeviceFromName = name => {
    return this.state.devices.find(e => e.name === name);
  };

  render() {
    const { devices, loader } = this.state;

    const Main = (
      <div className="App">
        <Header />
        <div className="wrapper">
          <Route
            exact
            path="/"
            render={() => <DeviceList devices={devices} />}
          />
          <Route
            path="/device/:name"
            render={props => (
              <DeviceDetails
                {...props}
                getDeviceFromName={this.getDeviceFromName}
              />
            )}
          />
        </div>
      </div>
    );

    return loader ? <Loader /> : Main;
  }
}

export default App;
