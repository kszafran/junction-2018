import React, { Component } from "react";
import DisplayMode from "../DisplayMode";
import HistoryChart from "../HistoryChart";

export default class DeviceDetails extends Component {
  constructor(props) {
    super(props);
    const name = this.props.match.params.name;
    const device = this.props.getDeviceFromName(name);

    const inititalSelect = device ? Object.keys(device.data_history)[0] : ""; // URGH

    this.state = {
      selectedId: inititalSelect // URGHHHH
    };
  }

  changeSelectedId = event => {
    const value = event.target.value;

    this.setState({ selectedId: value });
  };

  render() {
    const name = this.props.match.params.name;
    const device = this.props.getDeviceFromName(name);
    const chartData = {
      [this.state.selectedId]: device
        ? device.data_history[this.state.selectedId]
        : []
    };

    return (
      <React.Fragment>
        <HistoryChart
          data={chartData}
          selectedId={this.state.selectedId}
          name={name}
        />

        <DisplayMode
          device={device}
          selectedId={this.state.selectedId}
          changeSelectedId={this.changeSelectedId}
        />
      </React.Fragment>
    );
  }
}
