import React from "react";
import { Link } from "react-router-dom";
import "./index.css";

const Device = ({ data }) => {
  const { name, cisco_health_data } = data;
  console.log("DATA", data);

  const currentHealthData = cisco_health_data
    ? cisco_health_data.healthScore
    : 0;

  let healthColor;
  if (currentHealthData <= 3) {
    healthColor = "red-color";
  } else if (currentHealthData <= 5) {
    healthColor = "orange-color";
  } else if (currentHealthData <= 7) {
    healthColor = "yellow-color";
  } else {
    healthColor = "green-color";
  }

  return (
    <div>
      <div className="device">
        <div className="name">{name}</div>
        <div className={`status ${healthColor}`}>{currentHealthData}</div>

        <Link to={`/device/${name}`}>
          <button className="detail">Details</button>
        </Link>
      </div>
    </div>
  );
};

export default Device;
