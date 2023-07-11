import { useState, useEffect } from 'react';
import reactLogo from './assets/react.svg';
import Exam from './components/Exam';
import Register from './components/Register';
import Cookies from 'js-cookie';

function App() {
  const [name, setName] = useState(null);
  const [exams, setExams] = useState([]);

  useEffect(() => {
    const storedName = Cookies.get('name');
    if (storedName) {
      setName(storedName);
    }
  }, []);

  useEffect(() => {
    fetchExams();
  }, []);

  const fetchExams = async () => {
    try {
      const response = await fetch('http://localhost:3000/exams');
      const data = await response.json();
      setExams(data);
    } catch (error) {
      console.log('Error fetching exams:', error);
    }
  };

  const handleLogout = () => {
    Cookies.remove('name');
    setName(null);
  };

  const handleExamUpdate = () => {
    fetchExams();
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-800 to-gray-900 text-white">
      <header className="bg-gray-900 py-4">
        <div className="container mx-auto flex justify-between items-center">
          <h1 className="text-2xl font-bold">Exam Pro Mock Exam Scheduler: Bouman 6</h1>
          {name && (
            <button
              className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded"
              onClick={handleLogout}
            >
              Logout
            </button>
          )}
          <img src={reactLogo} alt="React Logo" className="w-8 h-8" />
        </div>
      </header>
      <main className="container mx-auto py-8">
        {name ? (
          <>
            <h2 className="text-2xl font-bold mb-4 text-white">Welcome, {name}!</h2>
            {exams.map((exam) => (
              <Exam key={exam.id} {...exam} userName={name} onUpdate={handleExamUpdate} />
            ))}
          </>
        ) : (
          <Register setName={setName} />
        )}
      </main>
    </div>
  );
}

export default App;
