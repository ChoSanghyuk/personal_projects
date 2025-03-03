import 'package:flutter/material.dart';
import '../presentation/funds.dart';
import 'funds_api.dart';



class FundsApiHttpMock implements FundsApi {
  FundsApiHttpMock();
  
  static bool _useAlternativeData = false;  // Add switch flag

  static void toggleDataSet() {
    _useAlternativeData = !_useAlternativeData;
  }

  Future<List<FundData>> getFundsData(int fundId) async {
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

  Future<List<FundTableData>> getFundsTableData(int fundId) async {
  
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
