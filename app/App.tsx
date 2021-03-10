import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { SafeAreaView, View, FlatList, StyleSheet, Text, TextInput } from 'react-native';

const DATA = [
    {
        id: 'bd7acbea-c1b1-46c2-aed5-3ad53abb28ba',
        name: 'duolingo',
        icon: '📚',
        count: 185,
    },
    {
        id: '3ac68afc-c605-48d3-a4f8-fbd91aa97f63',
        icon: '💪',
        name: '10 pressups everyday',
        count: 2,
    },
    {
        id: '58694a0f-3da1-471f-bd96-145571e29d72',
        name: 'develop app',
        icon: '💻',
        count: 5,
    },
];

const Item = ({ name, icon, count }) => (
    <View style={styles.item}>
        <View style={{ flex: 0.1 }}>
            <Text>{icon}</Text>
        </View>
        <View style={{ flex: 0.7 }}>
            <Text style={styles.name}>{name}</Text>
        </View>
        <View style={{ flex: 0.2 }}>
            <Text>🔥 {count}</Text>
        </View>
    </View>

);

const App = () => {
    const renderItem = ({ item }) => (
        <Item name={item.name} icon={item.icon} count={item.count} />
    );

    return (
        <SafeAreaView style={styles.container}>
            <Text style={{ fontSize: 28, padding: 20 }}>Streakify</Text>
                        
            <TextInput
                style={{ height: 80, padding: 20, fontSize: 18 }}
                placeholder="add new streak..."
            />

            <FlatList
                data={DATA}
                renderItem={renderItem}
                keyExtractor={item => item.id}
            />


            <StatusBar style="auto" />
        </SafeAreaView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        marginTop: StatusBar.currentHeight || 0,
    },
    item: {
        flexDirection: "row",
        backgroundColor: '#F8F8F8',
        padding: 20,
        marginVertical: 8,
        marginHorizontal: 16,
    },
    name: {
        fontSize: 20,
    },
});

export default App;