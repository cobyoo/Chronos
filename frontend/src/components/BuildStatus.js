import React, { useState, useEffect } from 'react';

function BuildStatus() {
    const [status, setStatus] = useState('Pending');

    useEffect(() => {
        fetch('/api/build-status')
            .then(response => response.json())
            .then(data => setStatus(data.status))
            .catch(error => console.error('Error fetching build status:', error));
    }, []);

    return (
        <div>
            <h2>Current Build Status: {status}</h2>
        </div>
    );
}

export default BuildStatus;
