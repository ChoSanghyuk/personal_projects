import 'dart:convert';
import 'package:http/http.dart' as http;
import '../presentation/funds.dart';
import './config_loader.dart';
import 'package:flutter/material.dart';


abstract class FundsApi {
  Future<List<FundData>> getFundsData(int index) ;
  Future<List<FundTableData>> getFundsTableData(int index);
}

class FundsApiProvider {
  static FundsApi getApi() {
    if (ConfigLoader.useMock()) {
      return FundsApiHttpMock();
    } else {
      return FundsApiHttp();
    }
  }
}


class FundsApiHttp implements FundsApi {
  FundsApiHttp();

  // New method to get funds data
  Future<List<FundData>> getFundsData(int index) async {
    try {
      final url = ConfigLoader.getUrl();
      final response = await http.get(Uri.parse('$url/funds/$index/portion'));
      
      if (response.statusCode == 200) {
        final Map<String, dynamic> jsonMap = json.decode(utf8.decode(response.bodyBytes));
        return [
          FundData('안전 자산', jsonMap['stable'].toDouble(), Colors.orange),
          FundData('변동 자산', jsonMap['volatile'].toDouble(), Colors.purple),
        ];
      } else {
        throw Exception('Failed to load funds data: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load funds data: $e');
    }
  }

  // New method to get funds table data
  Future<List<FundTableData>> getFundsTableData(int index) async {
    try {
      final url = ConfigLoader.getUrl();
      final response = await http.get(Uri.parse('$url/funds/$index/assets'));
      
      if (response.statusCode == 200) {
        final List<dynamic> jsonList = json.decode(utf8.decode(response.bodyBytes));
        return jsonList.map((json) => FundTableData(
          name: json['name'] ?? "",
          amount: json['amount'] ?? "",
          amountDollar: json['amount_dollar'] == '' ? '0' : json['amount_dollar'],
          profitRate: json['profitRate'] ?? "",
          division: json['division'] ?? "",
          quantity: json['quantity'] ?? "0",
          price: json['price'] == '' ? '0' : json['price'],
          priceDollar: json['price_dollar'] == '' ? '0' : json['price_dollar'],
          isStable: json['isStable'] ?? true,
        )).toList();
      } else {
        throw Exception('Failed to load funds table data: ${response.statusCode}');
      }
    } catch (e) {
      throw Exception('Failed to load funds table data: $e');
    }
  }
}



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
