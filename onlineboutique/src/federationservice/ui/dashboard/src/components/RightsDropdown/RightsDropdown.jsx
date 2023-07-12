import { RightsTagTrue, RightsTagFalse } from "./RightsDropdown.styles";
import { useState } from "react";

const rightsMapping = {
  canEarnCommission: "Earn Commission",
  canShareInventory: "Share Inventory",
  canShareData: "Share Data",
  canCoPromote: "Co-Promote",
  canSell: "Sell",
};

const RightsChip = ({ option, selected, onClick }) => {
  const displayText = rightsMapping[option];
  return selected ? (
    <RightsTagTrue onClick={onClick}>{displayText}</RightsTagTrue>
  ) : (
    <RightsTagFalse onClick={onClick}>{displayText}</RightsTagFalse>
  );
};
  
const RightsDropdown = ({ options, setSelectedRights, selectedRights }) => {
  const handleClick = (option) => {
    setSelectedRights(prevRights => ({
      ...prevRights,
      [option]: !prevRights[option]
    }));
  };

  return (
    <div>
      {options.map(option => (
        <RightsChip
          key={option}
          option={option}
          selected={selectedRights[option]}
          onClick={() => handleClick(option)}
        />
      ))}
    </div>
  );
};

export default RightsDropdown;
  