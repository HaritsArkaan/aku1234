import React, { useEffect, useState } from "react";
import axios from "axios";

function Dashboard() {
  const [input, setInput] = useState({
    judul: "",
    subJudul: "",
    kategori: "",
    deadline: "",
    deskripsi: "",
  });

  const [students, setStudents] = useState([]);

  // Mengambil data siswa
  useEffect(() => {
    fetchStudents();
  }, []);

  const fetchStudents = async () => {
    try {
      const res = await axios.get("http://localhost:8080/tugaspendahuluans");
      setStudents(res.data.data); // Sesuaikan dengan struktur respon dari backend
    } catch (err) {
      console.error("Error fetching data:", err);
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
      const studentData = {
        judul: input.judul,
        subJudul: input.subJudul,
        kategori: input.kategori,
        deadline: input.deadline,
        deskripsi: input.deskripsi,
      };
      console.log("Adding tugas:", studentData);
      // Mengirim data tugas baru
      const res = await axios.post(
        "http://localhost:8080/tugaspendahuluans",
        studentData
      );
      console.log("Added tugas:", res.data.data); // Log data yang ditambahkan
      // Menambahkan tugas baru ke dalam daftar tugas secara langsung
      setStudents((prevStudents) => [...prevStudents, res.data.data]);
      setInput({ judul: "", sub_judul: "", kategori: "", deadline: "", deskripsi: "", });
    } catch (err) {
      console.error(
        "Error adding tugas:",
        err.response ? err.response.data : err.message
      );
    }
  };

  // Menghapus siswa
  const handleDelete = (id) => {
    try {
      axios.delete(`http://localhost:8080/tugaspendahuluans/${id}`);
      fetchStudents(); // Refresh data setelah penghapusan
    } catch (err) {
      console.error("Error deleting student:", err);
    }
  };

  console.log(students);

  return (
    <div className="p-6 bg-gray-100 min-h-screen">
      <h1 className="text-3xl font-bold mb-4">Tugas Pendahuluan</h1>
      <div className="flex space-x-4 mb-4">
        <input
          type="text"
          name="judul"
          value={input.judul}
          onChange={handleInput}
          placeholder="Judul"
          className="border border-gray-300 rounded-md px-4 py-2"
        />
        <input
          type="text"
          name="subJudul"
          value={input.subJudul}
          onChange={handleInput}
          placeholder="subJudul"
          className="border border-gray-300 rounded-md px-4 py-2"
        />
        <input
          type="text"
          name="kategori"
          value={input.kategori}
          onChange={handleInput}
          placeholder="kategori"
          className="border border-gray-300 rounded-md px-4 py-2"
        />
        <input
          type="text"
          name="deadline"
          value={input.deadline}
          onChange={handleInput}
          placeholder="deadline"
          className="border border-gray-300 rounded-md px-4 py-2"
        />
        <input
          type="text"
          name="deskripsi"
          value={input.deskripsi}
          onChange={handleInput}
          placeholder="deskripsi"
          className="border border-gray-300 rounded-md px-4 py-2"
        />
        <button
          onClick={handleAddStudent}
          className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600"
        >
          Add Tugas
        </button>
      </div>
      <table className="table-auto w-full bg-white shadow-md rounded-lg">
        <thead>
          <tr className="bg-gray-100">
            <th className="px-4 py-2 text-left">Name</th>
            <th className="px-4 py-2 text-left">Sub Judul</th>
            <th className="px-4 py-2 text-left">Kategori</th>
            <th className="px-4 py-2 text-left">Deadline</th>
            <th className="px-4 py-2 text-left">Deskripsi</th>
            <th className="px-4 py-2 text-left">Actions</th>
          </tr>
        </thead>
        <tbody>
          {students.map((student) => (
            <tr key={student.ID} className="border-t">
              <td className="px-4 py-2">{student.judul}</td>
              <td className="px-4 py-2">{student.suJudul}</td>
              <td className="px-4 py-2">{student.kategori}</td>
              <td className="px-4 py-2">{student.deadline}</td>
              <td className="px-4 py-2">{student.deskripsi}</td>
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
