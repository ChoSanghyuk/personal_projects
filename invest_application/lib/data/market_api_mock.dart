class MarketApiMock {
  Future<Map<String, dynamic>> getFearGreedIndex() async {
    // Simulate network delay
    await Future.delayed(const Duration(milliseconds: 800));
    return {
      'value': 75,
      'status': 'Greed',
      'weeklyData': [65, 70, 72, 68, 73, 75, 75], // Mock weekly data
    };
  }

  Future<Map<String, dynamic>> getNasdaqData() async {
    await Future.delayed(const Duration(milliseconds: 800));
    return {
      'value': 15055.65,
      'change': 0.85,
      'weeklyData': [14800, 14900, 15100, 14950, 15000, 15055.65], // Mock weekly data
    };
  }

  Future<Map<String, dynamic>> getSP500Data() async {
    await Future.delayed(const Duration(milliseconds: 800));
    return {
      'value': 4783.35,
      'change': 1.2,
    };
  }
}
