import React from "react";
import "./index.css";

const DeviceDetails = props => {
  const id = parseInt(props.match.params.id, 10);
  const device = props.getDeviceFromID(id);
  return <div />;
};

export default DeviceDetails;
