import React from 'react';
import Cookies from 'js-cookie';

const Register = ({ setName }) => {
    const [localName, setLocalName] = React.useState(''); // Local state for the name input

    const handleFormSubmit = (event) => {
        event.preventDefault();
        Cookies.set('name', localName, { expires: 7 }); // Store the name in a cookie with a 7-day expiry
        setName(localName.trim());
    };

    const handleNameChange = (event) => {
        setLocalName(event.target.value); // Update the name state using the setName prop
    };

    return (
        <div>
            <h2 className="text-xl font-semibold">What is your name?</h2>
            <form onSubmit={handleFormSubmit}>
                <input
                    type="text"
                    value={localName}
                    onChange={handleNameChange}
                    className="border border-gray-300 rounded px-3 py-2 mt-2 text-gray-800 mr-2"
                    placeholder="Enter your name"
                />
                <button
                    type="submit"
                    className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-6 rounded mt-2"
                >
                    Save
                </button>
            </form>
        </div>
    );
};

export default Register;
