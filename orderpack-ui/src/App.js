import { useState } from "react";
import axios from "axios";
import "./App.css"; // Import CSS

const API_BASE_URL = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080";

function App() {
  const [packs, setPacks] = useState([]);
  const [quantity, setQuantity] = useState("");
  const [orderResult, setOrderResult] = useState(null);

  // Fetch pack sizes
  const fetchPacks = async () => {
    try {
      console.log("Fetching packs from API...");
      const response = await axios.get(`${API_BASE_URL}/packs`);
      console.log("API Response:", response.data);
      setPacks(response.data.packs);
    } catch (error) {
      console.error("Error fetching packs:", error);
    }
  };

  // Update pack sizes
  const updatePacks = async () => {
    try {
      const newPacks = prompt("Enter new pack sizes (comma separated):");
      if (!newPacks) return;
      const packArray = newPacks.split(",").map(Number);
      await axios.post(`${API_BASE_URL}/packs`, packArray);
      fetchPacks();
    } catch (error) {
      console.error("Error updating packs:", error);
    }
  };

  // Calculate best pack combination
  const calculateOrder = async () => {
    if (!quantity) return;
    try {
      const response = await axios.get(`${API_BASE_URL}/calculate?quantity=${quantity}`);
      setOrderResult(response.data);
    } catch (error) {
      console.error("Error calculating order:", error);
    }
  };

  // Clear the UI state
  const clearUI = () => {
    setPacks([]);
    setQuantity("");
    setOrderResult(null);
  };

  return (
      <div className="container">
        <h1>📦 OrderPack Optimizer</h1>

        <div className="card">
          <h2>Available Packs</h2>
          <p>{packs.length ? packs.join(", ") : "No packs available"}</p>
          <button onClick={updatePacks} className="btn">Update Packs</button>
        </div>

        <div className="card">
          <h2>Enter Order Quantity</h2>
          <input
              type="number"
              placeholder="Enter quantity"
              value={quantity}
              onChange={(e) => setQuantity(e.target.value)}
              className="input"
          />
          <button onClick={calculateOrder} className="btn">Calculate</button>
        </div>

        {orderResult && (
            <div className="card">
              <h2>Optimal Packs</h2>
              <p>Order Quantity: <strong>{orderResult.orderQuantity}</strong></p>
              {orderResult.packs.length > 0 ? (
                  <ul>
                    {orderResult.packs.map((pack, index) => (
                        <li key={index}>📦 {pack.packSize} x {pack.quantity}</li>
                    ))}
                  </ul>
              ) : (
                  <p>No packs available</p>
              )}
            </div>
        )}

        {/* Clear Button */}
        <button onClick={clearUI} className="btn-clear">Clear</button>
      </div>
  );
}

export default App;
