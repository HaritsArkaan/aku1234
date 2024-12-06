import React, { useEffect, useState } from "react";
import axios from "axios";

function Dashboard() {
  const [input, setInput] = useState({
    name: "",
    age: 0,
    class_id: 0,
  });

  const [students, setStudents] = useState([]);

  // Mengambil data siswa
  useEffect(() => {
    fetchStudents();
  }, []);

  const fetchStudents = async () => {
    try {
      const res = await axios.get("http://localhost:8080/students");
      setStudents(res.data.data); // Sesuaikan dengan struktur respon dari backend
    } catch (err) {
      console.error("Error fetching students:", err);
    }
  };

  // Menangani input form
  const handleInput = (e) => {
    const { name, value } = e.target;
    setInput({ ...input, [name]: value });
  };

  // Menambahkan siswa baru
  const handleAddStudent = async (e) => {
    e.preventDefault();
    try {
      // Pastikan age adalah integer
      const studentData = {
        name: input.name,
        age: parseInt(input.age, 10), // Mengonversi age ke integer
        class: parseInt(input.class_id, 10),
      };
      console.log("Adding student:", studentData);
      // Mengirim data siswa baru
      const res = await axios.post(
        "http://localhost:8080/students",
        studentData
      );
      console.log("Added student:", res.data.data); // Log data yang ditambahkan
      // Menambahkan siswa baru ke dalam daftar siswa secara langsung
      setStudents((prevStudents) => [...prevStudents, res.data.data]);
      setInput({ name: "", age: 0, class_id: 0 }); // Reset input
    } catch (err) {
      console.error(
        "Error adding student:",
        err.response ? err.response.data : err.message
      );
    }
  };

  // Menghapus siswa
  const handleDelete = (id) => {
    try {
      axios.delete(`http://localhost:8080/students/${id}`);
      fetchStudents(); // Refresh data setelah penghapusan
    } catch (err) {
      console.error("Error deleting student:", err);
    }
  };

  console.log(students);

  return (
    <div className="p-6 bg-gray-100 min-h-screen">
      <h1 className="text-3xl font-bold mb-4">Students</h1>
      <div className="flex space-x-4 mb-4">
        <input
          type="text"
          name="name"
          value={input.name}
          onChange={handleInput}
          placeholder="Name"
          className="border border-gray-300 rounded-md px-4 py-2"
        />
        <input
          type="number"
          name="age"
          value={input.age}
          onChange={handleInput}
          placeholder="Age"
          className="border border-gray-300 rounded-md px-4 py-2"
        />
        <input
          type="number"
          name="class_id"
          value={input.class_id}
          onChange={handleInput}
          placeholder="Class_id"
          className="border border-gray-300 rounded-md px-4 py-2"
        />
        <button
          onClick={handleAddStudent}
          className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
        >
          Add Student
        </button>
      </div>
      <table className="table-auto w-full bg-white shadow-md rounded-lg">
        <thead>
          <tr className="bg-gray-100">
            <th className="px-4 py-2 text-left">Name</th>
            <th className="px-4 py-2 text-left">Age</th>
            <th className="px-4 py-2 text-left">Class</th>
            <th className="px-4 py-2 text-left">Actions</th>
          </tr>
        </thead>
        <tbody>
          {students.map((student) => (
            <tr key={student.ID} className="border-t">
              <td className="px-4 py-2">{student.name}</td>
              <td className="px-4 py-2">{student.age}</td>
              <td className="px-4 py-2">{student.class_id}</td>
              <td className="px-4 py-2">
                <button className="bg-yellow-400 text-black px-3 py-1 rounded-md mr-2 hover:bg-yellow-500">
                  Edit
                </button>
                <button
                  type="submit"
                  onClick={() => handleDelete(student.ID)}
                  className="bg-red-500 text-white px-3 py-1 rounded-md hover:bg-red-600"
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default Dashboard;
