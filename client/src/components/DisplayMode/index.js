import React, { Component } from "react";
import RadioButton from "../RadioButton";

import "./index.css";

class DisplayModeButtons extends Component {
  render() {
    const { device, selectedId, changeSelectedId } = this.props;
    const dataHistory = device ? device.data_history : [];

    const dataHistoryButtons = Object.keys(dataHistory).map((e, i) => {
      return (
        <RadioButton
          id={e}
          selectedId={selectedId}
          changeSelectedId={changeSelectedId}
          key={i}
        />
      );
    });

    return (
      <div className="buttons">
        <div className="custom-radios">{dataHistoryButtons}</div>
      </div>
    );
  }
}

export default DisplayModeButtons;
