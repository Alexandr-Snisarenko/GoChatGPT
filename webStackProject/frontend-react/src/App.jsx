
import { useEffect, useState } from 'react'

function App() {
  const [users, setUsers] = useState([])
  const [name, setName] = useState('')
  const [email, setEmail] = useState('')

  useEffect(() => {
    fetch('/api/users')
      .then(res => res.json())
      .then(data => setUsers(data))
  }, [])

  const createUser = async () => {
    await fetch('/api/users', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ name, email: email || null })
    })
    setName('')
    setEmail('')
    const res = await fetch('/api/users')
    setUsers(await res.json())
  }

  const deleteUser = async (id) => {
    await fetch('/api/users/' + id, { method: 'DELETE' })
    setUsers(users.filter(u => u.id !== id))
  }

  return (
    <div style={{ padding: '2rem' }}>
      <h1>Users</h1>
      <input
        placeholder="Name"
        value={name}
        onChange={e => setName(e.target.value)}
      />
      <input
        placeholder="Email"
        value={email}
        onChange={e => setEmail(e.target.value)}
      />
      <button onClick={createUser}>Create</button>
      <ul>
        {users.map(u => (
          <li key={u.id}>
            {u.name} ({u.email || 'No email'})
            <button onClick={() => deleteUser(u.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default App
