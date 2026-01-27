import Navbar from "../components/Navbar";
import ProductCard from "../components/ProductCard";
import { dummyProducts } from "../data/product";

function App() {
  return (
    <div className="bg-gray-100 min-h-screen">
      <Navbar />

      <main className="container mx-auto p-4 sm:p-8">
        <h1 className="text-3xl font-bold text-gray-900 mb-6">
          Featured Products
        </h1>

        <div
          className="
          grid grid-cols-1 
          gap-6 
          sm:grid-cols-2 
          lg:grid-cols-3 
          xl:grid-cols-4
        "
        >
          {dummyProducts.map((product) => (
            <ProductCard key={product.id} product={product} />
          ))}
        </div>
      </main>
    </div>
  );
}

export default App;
