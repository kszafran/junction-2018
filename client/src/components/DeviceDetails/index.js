import React from "react";
import "./index.css";

const DeviceDetails = props => {
  const id = props.match.params.id;
  const device = props.getDeviceFromID(id);

  return <div />;
};

export default DeviceDetails;
