import React, { Component } from "react";
import "./index.css";
import Device from "../Device";

const MOCK_DATA_URL =
  "https://gist.githubusercontent.com/vietdang7/a5d495df7bc398a1f6c46eade07d452d/raw/d4f7639672d99e1556300c4404356d436e64e854/Junction%2520Data";

export default class DeviceList extends Component {
  constructor(props) {
    super(props);
    this.state = { devices: [] };
  }

  componentDidMount = async () => {
    const result = await fetch(MOCK_DATA_URL);
    const devices = await result.json();

    this.setState({
      devices
    });
  };

  render() {
    const { devices } = this.state;

    return (
      <React.Fragment>
        <div className="title">Device list</div>
        <div className="list-title">
          <div className="name">Device Name</div>
          <div className="status">Status</div>
        </div>
        <div className="device-list">
          {devices.map((e, i) => (
            <Device data={e} key={i} />
          ))}
        </div>
      </React.Fragment>
    );
  }
}
