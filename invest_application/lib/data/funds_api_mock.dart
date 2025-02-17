import 'package:flutter/material.dart';
import '../presentation/funds.dart';

class FundData {
  final String name;
  final double value;
  final Color color;

  FundData(this.name, this.value, this.color);
}

class FundsApiMock {
  static bool _useAlternativeData = false;  // Add switch flag

  static void toggleDataSet() {
    _useAlternativeData = !_useAlternativeData;
  }

  static List<FundData> getFundsData() {
    toggleDataSet();
    if (_useAlternativeData) {
      return [
        FundData('안전 자산', 20.0, Colors.orange),
        FundData('변동 자산', 80.0, Colors.purple),
      ];
    }
    return [
      FundData('안전 자산', 80.0, Colors.orange),
      FundData('변동 자산', 20.0, Colors.purple),
    ];
  }

  static List<FundTableData> getFundsTableData() {
    return [
      FundTableData(
        name: '금',
        amount: '130000000.0',
        amountDollar: '100000.0',
        profitRate: '5.2',
        division: 'Bonds',
        quantity: '100.0',
        price: '1300000.0',
        priceDollar: '1000.0',
        isStable: true,
      ),
      FundTableData(
        name: '삼성 전자',
        amount: '130000000.0',
        amountDollar: '100000.0',
        profitRate: '8.7',
        division: 'Stocks',
        quantity: '50.0',
        price: '2600000.0',
        priceDollar: '2000.0',
        isStable: false,
      ),
      FundTableData(
        name: '원화',
        amount: '50000.0',
        amountDollar: '0',
        profitRate: '2.1',
        division: 'Cash',
        quantity: '1.0',
        price: '50000.0',
        priceDollar: '0',
        isStable: true,
      ),
      FundTableData(
        name: '비트코인',
        amount: '75000.0',
        amountDollar: '75000.0',
        profitRate: '6.5',
        division: 'Alternative',
        quantity: '25.0',
        price: '3000.0',
        priceDollar: '3000.0',
        isStable: false,
      ),
    ];
  }
}
