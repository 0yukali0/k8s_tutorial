import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { Button } from 'antd';

const fetchData = async () => {
  const response = await axios.get('http://localhost:8080/res');
  return response.data;
};

export default function App() {
  const { data, isLoading, isError, refetch } = useQuery({
    queryKey: ['machineInfo'],
    queryFn: fetchData,
  });

  if (isLoading) return <div>Loading...</div>;
  if (isError) return <div>Error: {data?.message}</div>;

  return (
    <div>
      <h1>Hello</h1>
      <h2>Machine Info:</h2>
      {data && (
        <div>
          <p>CPU Usage: {data.cpu.toFixed(2)}%</p>
          <p>Memory Usage: {data.memory.toFixed(2)}%</p>
        </div>
      )}
      <Button onClick={() => refetch()}>Refresh</Button>
    </div>
  );
}
