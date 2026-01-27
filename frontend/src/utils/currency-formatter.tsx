import React from "react";

interface CurrencyFormatterProps {
  amount: number;
  className?: string;
}

const CurrencyFormatter: React.FC<CurrencyFormatterProps> = ({
  amount,
  className,
}) => {
  const formatter = new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  });

  const formattedAmount = formatter.format(amount);

  return <p className={className}>{formattedAmount}</p>;
};

export default CurrencyFormatter;
