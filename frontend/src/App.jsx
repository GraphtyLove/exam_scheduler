import { useState, useEffect } from 'react';
import { BeatLoader } from 'react-spinners';
import Exam from './components/Exam';
import Register from './components/Register';
import Cookies from 'js-cookie';
import { API_URL } from './constants';

function App() {
  const [name, setName] = useState(null);
  const [exams, setExams] = useState([]);
  const [loading, setLoading] = useState(false);

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
    setLoading(true);
    try {
      const response = await fetch(`${API_URL}/exams`);
      const data = await response.json();
      setExams(data);
    } catch (error) {
      console.log('Error fetching exams:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    Cookies.remove('name');
    setName(null);
  };

  const handleExamUpdate = () => {
    fetchExams();
  };

  const capitalizeFirstLetter = string => {
    return string.charAt(0).toUpperCase() + string.slice(1);
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-800 to-gray-900 text-white">
      <header className="bg-gray-900 py-4">
        <div className="container mx-auto flex justify-center items-center">
          <div className="flex">
            <img src="/becode.png" alt="Becode Logo" className="w-12 h-10 mr-5" />
            <h1 className="text-2xl font-bold">ExamPro Exam Scheduler</h1>
          </div>

          {name && (
            <div className="flex items-center text-2xl ml-auto">
              <p className="mr-5">Welcome, {capitalizeFirstLetter(name.toLowerCase())}!</p>
              <button
                className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded"
                onClick={handleLogout}
              >
                Logout
              </button>
            </div>
          )}
        </div>
      </header>
      <main className="container mx-auto py-8">
        {name ? (
          <>
            {loading ? (
              <div className="flex justify-center items-center">
                <BeatLoader color={'#ffffff'} loading={loading} size={15} />
              </div>
            ) : (
              exams.map((exam) => (
                <Exam key={exam.id} {...exam} userName={name} onUpdate={handleExamUpdate} />
              ))
            )}
          </>
        ) : (
          <Register setName={setName} />
        )}
      </main>
    </div>
  );
}

export default App;
