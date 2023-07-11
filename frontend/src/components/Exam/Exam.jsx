import React from 'react';
import { useState, useEffect } from 'react';


const Exam = ({ id, name: examName, isInProgress, challenger, startTime, endTime, userName, onUpdate }) => {
    const [timeLeft, setTimeLeft] = useState('Exam Ended');

    useEffect(() => {
        if (isInProgress) {
            const timer = setInterval(() => {
                setTimeLeft(calculateTimeLeft());
            }, 1000);

            return () => {
                clearInterval(timer);
            };
        }
    }, [isInProgress, endTime]);

    const handleButtonClick = async () => {
        if (isInProgress) {
            // Logic for finishing the exam
            console.log('Finish exam clicked');

            const updatedExam = {
                id: id,
                name: examName,
                isInProgress: false,
                challenger: "",
                startTime: null,
                endTime: null,
            };

            try {
                const response = await fetch(`http://localhost:3000/exam/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(updatedExam),
                });

                if (response.ok) {
                    console.log('Exam finished successfully');
                    onUpdate(); // Fetch updated data from API
                    // Perform any additional logic or state updates
                } else {
                    console.log('Failed to finish the exam');
                    // Handle the error case
                }
            } catch (error) {
                console.log('Error finishing the exam:', error);
                // Handle the error case
            }
        } else {
            // Logic for booking the exam
            console.log('Book this exam clicked');

            const updatedExam = {
                id: id,
                name: examName,
                isInProgress: true,
                challenger: userName,
                startTime: Date.now(),
                endTime: Date.now() + 90 * 60 * 1000, // Adding 90 minutes in milliseconds
            };

            try {
                const response = await fetch(`http://localhost:3000/exam/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(updatedExam),
                });

                if (response.ok) {
                    console.log('Exam booked successfully');
                    onUpdate(); // Fetch updated data from API
                    // Perform any additional logic or state updates
                } else {
                    console.log('Failed to book the exam');
                    // Handle the error case
                }
            } catch (error) {
                console.log('Error booking the exam:', error);
                // Handle the error case
            }
        }
    };

    const calculateTimeLeft = () => {
        const currentTime = Date.now();
        const timeLeft = endTime - currentTime;

        // Check if timeLeft is positive
        if (timeLeft > 0) {
            // Calculate hours, minutes, and seconds
            const hours = Math.floor(timeLeft / (1000 * 60 * 60));
            const minutes = Math.floor((timeLeft % (1000 * 60 * 60)) / (1000 * 60));
            const seconds = Math.floor((timeLeft % (1000 * 60)) / 1000);

            return `${hours}h ${minutes}m ${seconds}s`;
        }

        return 'Exam Ended'; // Display "Exam Ended" when timeLeft is zero or negative
    };


    return (
        <div className="max-w-md mx-auto bg-white p-4 rounded shadow my-5">
            <div className="mb-4">
                <h2 className="text-xl font-semibold text-gray-800">{examName}</h2>
                <p className="text-gray-500">{isInProgress ? 'In Progress' : 'Not In Progress'}</p>
            </div>

            {isInProgress && (
                <>
                    <div className="mb-2">
                        <p className="font-medium text-gray-800">Challenger:</p>
                        <p className="text-gray-800">{challenger}</p>
                    </div>
                    <div className="mb-2">
                        <p className="font-medium text-gray-800">Time Left:</p>
                        <p className="font-bold text-red-500">{timeLeft}</p>
                    </div>
                </>
            )}

            <div className="mt-4">
                {isInProgress ? (
                    <button className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded" onClick={handleButtonClick}>
                        Finish Exam
                    </button>
                ) : (
                    <button className="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded" onClick={handleButtonClick}>
                        Book This Exam
                    </button>
                )}
            </div>
        </div>
    );
};

export default Exam;