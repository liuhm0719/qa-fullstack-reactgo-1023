import React, { useEffect, useState } from 'react';
import { createRoot } from 'react-dom/client';

// (a) bindings form: the API base is READ from the environment, never hardcoded.
// Uninjected -> fall back to same-origin relative (dev / co-located serving).
const API_BASE = (import.meta.env.VITE_API_BASE || '').replace(/\/$/, '');

function App() {
  const [msg, setMsg] = useState('loading…');
  const [err, setErr] = useState('');
  useEffect(() => {
    fetch(`${API_BASE}/api/hello`)
      .then(r => r.ok ? r.json() : Promise.reject(new Error('HTTP ' + r.status)))
      .then(d => setMsg(d.message || JSON.stringify(d)))
      .catch(e => { setErr(String(e)); setMsg('backend unreachable'); });
  }, []);
  return (
    <main style={{ fontFamily: 'sans-serif', padding: '2rem' }}>
      <h1>qa-fullstack reactgo</h1>
      <p>API_BASE={JSON.stringify(API_BASE)}</p>
      <p>backend says: <b>{msg}</b></p>
      {err && <p style={{ color: 'crimson' }}>error: {err}</p>}
    </main>
  );
}
createRoot(document.getElementById('root')).render(<App />);
