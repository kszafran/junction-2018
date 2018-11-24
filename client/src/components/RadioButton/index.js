import React from "react";

import "./index.css";

const FooterRadioButton = ({ id, selectedId, changeSelectedId }) => {
  return (
    <div className="category">
      <input
        type="radio"
        id={id}
        value={id}
        checked={selectedId === id}
        onChange={changeSelectedId}
      />
      <label htmlFor={id}>{id}</label>
    </div>
  );
};

export default FooterRadioButton;
