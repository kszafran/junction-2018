import React from "react";
import avatarPlaceholder from "../../assets/images/avatar.png";
import appLogo from "../../assets/images/logo_top.png";
import { Link } from "react-router-dom";

import "./index.css";

const Header = () => {
  return (
    <div className="header">
      <div />
      <Link to="/">
        <img src={appLogo} alt="app logo" />
      </Link>
      <img src={avatarPlaceholder} alt="avatar place holder" />
    </div>
  );
};

export default Header;
