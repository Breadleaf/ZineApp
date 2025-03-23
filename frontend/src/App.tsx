import React, { useEffect, useState } from "react";

interface Person {
  name: string;
  age: number;
}

interface Payload {
  people: Person[];
}

function App() {
  const [data, setData] = useState<Payload | null>(null);

  useEffect(() => {
    fetch("/api")
      .then((resp) => resp.json())
      .then(setData)
      .catch((err) => console.error("Error fetching data:", err));
  }, []);

  return (
    <div>
      <h1>Zine App</h1>
      {data ? (
        <ul>
          {data.people.map((person, index) => (
            <li key={index}>
              {person.name} - {person.age} years old
            </li>
          ))}
        </ul>
      ) : (
        <p>Loading data from backend...</p>
      )}
    </div>
  );
}

export default App;
