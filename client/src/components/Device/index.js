import React from "react";
import { Link } from "react-router-dom";
import "./index.css";

const Device = ({ data }) => {
  const {
    id,
    name,
    cisco_health_data: { current_health_data }
  } = data;
  console.log("id,", id);

  return (
    <div>
      <div className="device">
        <div className="name">{name}</div>
        <div className="status">{current_health_data}</div>

        <Link to={`device/${id}`}>
          <button className="detail">Details</button>
        </Link>
      </div>
    </div>
  );
};

export default Device;
